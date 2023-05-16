package api

// TODO: implement the api for the users

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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

func (a *API) GetListOfStocks(w http.ResponseWriter, r *http.Request) {
	a.enableCors(&w)

	apiKey := os.Getenv("TWELVE_API")
	key := r.URL.Query().Get("key")
	userId, err := strconv.ParseInt(r.URL.Query().Get("userId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cachedData, err := a.returnCacheData(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(cachedData) > 0 {
		type Stocks struct{}
		var stocks []Stocks
		err := json.Unmarshal(cachedData, &stocks)
		if err != nil {
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&stocks)
		return
	}

	stocks, err := a.DBStorage.GetStocksFromUser(userId)

	baseURL := "https://api.twelvedata.com/time_series?sysmbol="
	for _, stock := range stocks.([]database.Stock) {
		baseURL += stock.Symbol + ","
	}
	baseURL += "&interval=1day&apikey=" + apiKey + "source=docs"

	resp, err := http.Get(baseURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := a.insertOnCache(resp.Body, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Cache miss")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&resp.Body)
}
