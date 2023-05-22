package api

// TODO: implement the RSS feed

import (
	"encoding/json"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func NewRss() *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://br.investing.com/rss/stock_Fundamental.rss")
	return feed
}

func (api *API) HandleGetRssTitles(w http.ResponseWriter, r *http.Request) {

	feed := NewRss()

	type Rss struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Published   string `json:"published"`
		Image       string `json:"image"`
	}

	var rss []Rss = make([]Rss, 0)

	for _, item := range feed.Items {

		rss = append(rss, Rss{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			Author:      item.Author.Name,
			Published:   item.Published,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rss)
}
