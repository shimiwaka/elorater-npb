package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

type searchHandler struct{}

type SearchResponse struct {
	Players []SearchedPlayer `json:"players"`
}

type SearchedPlayer struct {
	Name string `json:"name"`
	Rate int    `json:"rate"`
	Id   uint   `json:"id"`
}

func search(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if db == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to db connection\"}")
		return
	}
	resp := SearchResponse{}
	players := []Player{}

	page, _ := strconv.Atoi(r.URL.Query().Get("p"))
	err := db.Limit(100).Offset(page * 100).Order("rate desc").Find(&players).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to fetch player data\"}")
		return
	}

	if len(players) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"incorrect specified page number\"}")
		return
	}

	for _, player := range players {
		resp.Players = append(resp.Players, SearchedPlayer{Name: player.Name, Rate: player.Rate, Id: player.ID})
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to encode json\"}")
		return
	}
	fmt.Fprint(w, buf.String())
}

func (p *searchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	search(db, w, r)
	db.Close()
}
