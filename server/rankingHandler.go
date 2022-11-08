package main

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"bytes"
)

type rankingHandler struct{}

type RankingResponse struct{
	Players []RankedPlayer	`json:"players"`
}

type RankedPlayer struct {
	Name string				`json:"name"`
	Rate int				`json:"rate"`
	Id uint					`json:"id"`
}

func (p *rankingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()

	resp := RankingResponse{}
	players := []Player{}

	page, _ := strconv.Atoi(r.URL.Query().Get("p"))
	db.Limit(100).Offset(page * 100).Order("rate desc").Find(&players)

	if len(players) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"incorrect specified page number\"}")
		return
	}

	for _, player := range players {
		resp.Players = append(resp.Players, RankedPlayer{Name: player.Name, Rate: player.Rate, Id: player.ID})
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