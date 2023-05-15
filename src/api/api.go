package api

// TODO: implement the api for the users

import (
	"encoding/json"
	"log"
	"net/http"

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

	data, err := a.cache.Get("daily")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(data.([]byte)) > 0 {
		type User struct {
			Id       int    `json:"id"`
			Username string `json:"username"`
		}
		var users []User
		err := json.Unmarshal(data.([]byte), &users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&users)
		return
	}

	log.Println("Cache miss")

	users, err := a.DBStorage.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := a.cache.Insert(b); err != nil {
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
