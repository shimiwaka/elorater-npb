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
}

func (p *rankingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
 
	db := ConnectDB()

	resp := RankingResponse{}

	players := []Player{}

	v := r.URL.Query()
	page, _ := strconv.Atoi(v.Get("p"))
	db.Limit(100).Offset(page * 100).Order("rate desc").Find(&players)

	for _, player := range players {
		// fmt.Fprintf(w, "%s:%d\n", player.Name, player.Rate)
		rPlayer := RankingPlayer{Name: player.Name, Rate: player.Rate}
		resp.Players = append(resp.Players, rPlayer)
	}
	
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&resp); err != nil {
		panic("encode error")
	}
	fmt.Fprint(w, buf.String())
}