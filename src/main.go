package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/victorguidi/market/api"
	"github.com/victorguidi/market/database"
)

func main() {

	db, err := database.NewDatabase("./databases/market.db")
	if err != nil {
		panic(err)
	}
	db.Init()

	cache := database.NewCache()
	cache.Init()

	api := api.NewAPI(":8080", db, cache)

	http.HandleFunc("/api/v1/users", api.GetUsers)
	http.HandleFunc("/api/v1/users/stocks/", api.GetListOfStocks)
	http.HandleFunc("/api/v1/stocks/", api.HandleGetOverviewStock)
	http.HandleFunc("/api/v1/post/user", api.PostUser)
	http.HandleFunc("/api/v1/post/stock", api.HandleInsertNewStock)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
