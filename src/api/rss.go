package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/mmcdole/gofeed"
)

type Rss struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	MainLink    string   `json:"mainLink"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Author      string   `json:"author"`
	Published   string   `json:"published"`
	Image       string   `json:"image"`
}

func (a *API) HandleGetRssTitles(w http.ResponseWriter, r *http.Request) {

	a.enableCors(&w)

	var rss []Rss = make([]Rss, 0)

	feed, err := a.DBStorage.GetAllLinkRss()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, item := range feed {
		fp := gofeed.NewParser()
		nfeed, err := fp.ParseURL(item)
		if err != nil {
			panic(err)
		}
		for _, item := range nfeed.Items {
			rss = append(rss, Rss{
				Title:       item.Title,
				Link:        item.Link,
				MainLink:    feed[i],
				Description: item.Description,
				Categories:  item.Categories,
				Author:      item.Author.Name,
				Published:   item.Published,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rss)
}

func (a *API) HandleGetRssTitlesFromUser(w http.ResponseWriter, r *http.Request) {
	a.enableCors(&w)

	userId, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[5], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var rss []Rss = make([]Rss, 0)

	feed, err := a.DBStorage.GetAllLinkRssForUser(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, item := range feed {
		fp := gofeed.NewParser()
		nfeed, err := fp.ParseURL(item)
		if err != nil {
			panic(err)
		}
		for _, item := range nfeed.Items {
			rss = append(rss, Rss{
				Title:       item.Title,
				Link:        item.Link,
				MainLink:    feed[i],
				Description: item.Description,
				Categories:  item.Categories,
				Author:      item.Author.Name,
				Published:   item.Published,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&rss)
}

func (a *API) HandleAddRssLink(w http.ResponseWriter, r *http.Request) {

	a.enableCors(&w)

	userId, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[5], 10, 64)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&req)
}
