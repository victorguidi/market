package api

import (
	"net/http"

	"github.com/victorguidi/market/database"
)

type API struct {
	listenAddr string
	Storage    database.Storage
}

func NewAPI(listenAddr string, store database.Storage) *API {
	return &API{
		listenAddr: listenAddr,
		Storage:    store,
	}
}

func (s *API) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (a *API) GetUsers(w http.ResponseWriter, r *http.Request) (value, err error) {
	return nil, nil
}

func (a *API) GetOneUserById(w http.ResponseWriter, r *http.Request) (*User, error) {
	return nil, nil
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
