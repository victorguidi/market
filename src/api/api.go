package api

// TODO: implement the api for the users

import (
	"encoding/json"
	"net/http"

	"github.com/victorguidi/market/database"
)

type API struct {
	listenAddr string
	DBStorage  database.DBStorage
}

func NewAPI(listenAddr string, store database.DBStorage) *API {
	return &API{
		listenAddr: listenAddr,
		DBStorage:  store,
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

	json.NewEncoder(w).Encode(&users)
}

func (a *API) GetOneUserById(w http.ResponseWriter, r *http.Request) {
}

func (a *API) Post(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) Put(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
