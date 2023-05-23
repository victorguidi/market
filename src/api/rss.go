package api

// TODO: implement the RSS feed

import (
	"encoding/json"
	"net/http"

	"github.com/mmcdole/gofeed"
)

type Feeds struct {
	Rss *gofeed.Feed
}

func (a *API) NewRss(links []string) *[]gofeed.Feed {
	feeds := make([]gofeed.Feed, len(links))
	for i, link := range links {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(link)
		if err != nil {
			panic(err)
		}
		feeds[i] = *feed
	}
	return &feeds
}

func (a *API) HandleGetRssTitles(w http.ResponseWriter, r *http.Request) {

	type Rss struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Published   string `json:"published"`
		Image       string `json:"image"`
	}

	var rss []Rss = make([]Rss, 0)

	for _, item := range a.feed {
		for _, feed := range item.Rss.Items {
			rss = append(rss, Rss{
				Title:       feed.Title,
				Link:        feed.Link,
				Description: feed.Description,
				Author:      feed.Author.Name,
				Published:   feed.Published,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rss)
}
