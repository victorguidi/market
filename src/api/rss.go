package api

// TODO: implement the RSS feed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mmcdole/gofeed"
)

type Feeds struct {
	Rss *gofeed.Feed
}

func (a *API) NewRss(links []string) error {
	if len(links) == 0 {
		return fmt.Errorf("no links to parse")
	}
	if len(a.feed) > 0 {
		a.feed = make([]Feeds, 0)
	}
	for _, link := range links {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(link)
		if err != nil {
			panic(err)
		}
		a.feed = append(a.feed, Feeds{Rss: feed})
	}
	return nil
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

	feed, err := a.DBStorage.GetAllLinkRss()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range feed {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(item)
		if err != nil {
			panic(err)
		}
		rss = append(rss, Rss{
			Title:       feed.Title,
			Link:        feed.Link,
			Description: feed.Description,
			Author:      feed.Author.Name,
			Published:   feed.Published,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rss)
}

func (a *API) HandleGetRssTitlesFromUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[4], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Rss struct {
		Title       string `json:"title"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Published   string `json:"published"`
		Image       string `json:"image"`
	}

	var rss []Rss = make([]Rss, 0)

	feed, err := a.DBStorage.GetAllLinkRssForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range feed {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(item)
		if err != nil {
			panic(err)
		}
		rss = append(rss, Rss{
			Title:       feed.Title,
			Link:        feed.Link,
			Description: feed.Description,
			Author:      feed.Author.Name,
			Published:   feed.Published,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rss)
}

func (a *API) HandleAddRssLink(w http.ResponseWriter, r *http.Request) {

	a.enableCors(&w)

	userId, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[4], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Req struct {
		Link string `json:"link"`
	}
	var req Req

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.DBStorage.InsertNewLinkRss(req.Link, userId)
	a.NewRss([]string{req.Link})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&req)
}
