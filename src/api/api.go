package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"strconv"

	"github.com/victorguidi/market/database"
)

type API struct {
	listenAddr string
	DBStorage  database.DBStorage
	cache      database.CStorage
}

func NewAPI(listenAddr string, store database.DBStorage, cache database.CStorage) *API {
	return &API{
		listenAddr: listenAddr,
		DBStorage:  store,
		cache:      cache,
	}
}

func (s *API) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (a *API) GetUsers(w http.ResponseWriter, r *http.Request) {
	a.enableCors(&w)

	users, err := a.DBStorage.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func (a *API) GetOneUserById(w http.ResponseWriter, r *http.Request) {
}

func (a *API) PostUser(w http.ResponseWriter, r *http.Request) {
	a.enableCors(&w)

	var user database.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.DBStorage.InsertUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func (a *API) Put(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) returnCacheData(key string) ([]byte, error) {
	data, err := a.cache.Get(key)
	if err != nil {
		return nil, err
	}

	if len(data.([]byte)) > 0 {
		return data.([]byte), nil
	}

	return nil, nil
}

func (a *API) insertOnCache(data any, key string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := a.cache.Insert(b, key); err != nil {
		return err
	}
	return nil
}

func (a *API) getFromExternalAPI(baseURL, key string) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	JSONBody := make(map[string]interface{})
	err = json.Unmarshal(body, &JSONBody)
	if err != nil {
		return nil, err
	}
	if JSONBody["status"] == "error" {
		return nil, err
	}

	if key != "skip" {
		if err := a.insertOnCache(JSONBody, key); err != nil {
			return nil, err
		}
	}

	return JSONBody, nil
}

func (a *API) GetListOfStocks(w http.ResponseWriter, r *http.Request) {
	a.enableCors(&w)

	apiKey := os.Getenv("TWELVE_API")

	key := strings.Split(r.URL.Path, "/")[5]
	userId, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[6], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cachedData, err := a.returnCacheData(key)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(cachedData) > 0 {
		JSONBody := make(map[string]interface{})
		err = json.Unmarshal(cachedData, &JSONBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&JSONBody)
		return
	}

	stocks, err := a.DBStorage.GetStocksFromUser(userId)
	log.Println(stocks)

	baseURL := "https://api.twelvedata.com/time_series?symbol="
	for i, stock := range stocks {
		if i == 7 {
			break
		}
		if i == len(stocks)-1 {
			baseURL += stock.Symbol
			continue
		}
		baseURL += stock.Symbol + ","
	}
	baseURL += "&interval=1day&apikey=" + apiKey + "&source=docs"

	body, err := a.getFromExternalAPI(baseURL, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Cache miss")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}

func (a *API) HandleInsertNewStock(w http.ResponseWriter, r *http.Request) {
	a.enableCors(&w)

	api := os.Getenv("ALPHA_API")

	type Req struct {
		Symbol string `json:"symbol"`
	}
	var req Req

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := "https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + req.Symbol + "&apikey=" + api

	body, err := a.getFromExternalAPI(url, "skip")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var stock database.Stock
	err = json.Unmarshal(jsonBytes, &stock)

	if err := a.DBStorage.InsertNewStockInfo(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&stock)
}

func (a *API) HandleGetOverviewStock(w http.ResponseWriter, r *http.Request) {

	a.enableCors(&w)

	api := os.Getenv("ALPHA_API")

	symbol := strings.Split(r.URL.Path, "/")[4]
	url := "https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + symbol + "&apikey=" + api

	check, err := a.DBStorage.GetStockAndCheckLastUpdate(symbol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if check {
		body, err := a.getFromExternalAPI(url, "overview")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var stock database.Stock
		err = json.Unmarshal(jsonBytes, &stock)

		if err := a.DBStorage.UpdateStockInfo(symbol, &stock); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&body)
		return
	}

	stock, err := a.DBStorage.GetStockInfo(symbol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&stock)

}

//TODO: Seach for a stock
