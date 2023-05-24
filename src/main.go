package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/victorguidi/market/api"
	"github.com/victorguidi/market/database"
)

func main() {

	// TODO: Fix RSS

	db, err := database.NewDatabase("./databases/market.db")
	if err != nil {
		panic(err)
	}
	db.Init()

	cache := database.NewCache()
	cache.Init()

	api := api.NewAPI(":8080", db, cache)

	// User API
	http.HandleFunc("/api/v1/users", api.GetUsers)
	http.HandleFunc("/api/v1/users/stocks/", api.GetListOfStocks)

	// Stock API
	http.HandleFunc("/api/v1/stocks/", api.HandleGetOverviewStock)
	http.HandleFunc("/api/v1/post/user", api.PostUser)
	http.HandleFunc("/api/v1/post/stock", api.HandleInsertNewStock)

	// RSS feed
	http.HandleFunc("/api/v1/rss/", api.HandleGetRssTitles)
	http.HandleFunc("/api/v1/rss/get", api.HandleAddRssLink)
	http.HandleFunc("/api/v1/rss/add", api.HandleAddRssLink)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
