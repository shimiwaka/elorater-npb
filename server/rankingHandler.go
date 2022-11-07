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
	Players []RankingPlayer	`json:"players"`
}

type RankingPlayer struct {
	Name string				`json:"name"`
	Rate int				`json:"rate"`
	Id uint					`json:"id"`
}

func (p *rankingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()

	resp := RankingResponse{}

	players := []Player{}

	v := r.URL.Query()
	page, _ := strconv.Atoi(v.Get("p"))
	db.Limit(100).Offset(page * 100).Order("rate desc").Find(&players)

	for _, player := range players {
		// fmt.Fprintf(w, "%s:%d\n", player.Name, player.Rate)
		rPlayer := RankingPlayer{Name: player.Name, Rate: player.Rate, Id: player.ID}
		resp.Players = append(resp.Players, rPlayer)
	}
	
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&resp); err != nil {
		panic("encode error")
	}
	fmt.Fprint(w, buf.String())
}