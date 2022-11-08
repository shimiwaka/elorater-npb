package main

import (
	"net/http"
	"fmt"
	"strconv"
)

type voteHandler struct{}

func (p *voteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var token Token
	var player1, player2 Player
	var c int
	var tokenString string

	vars := r.URL.Query()
	if vars == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to get query\"}")
		return
	}

	tokenString = vars.Get("token")
	c, _ = strconv.Atoi(vars.Get("c"))

	if c != 1 && c != 2 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"incorrect parameter\"}")
		return
	}

	db := ConnectDB()
	db.First(&token, "token = ?", tokenString)

	if token.Player1_id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"incorrect token\"}")
		return
	}

	db.First(&player1, token.Player1_id)
	db.First(&player2, token.Player2_id)
	
	if c == 1 {
		player1.Rate += int(32 * (((float32(player2.Rate) - float32(player1.Rate)) / 800) + 0.5))
		player2.Rate -= int(32 * (((float32(player2.Rate) - float32(player1.Rate)) / 800) + 0.5))
	} else {
		player1.Rate -= int(32 * (((float32(player2.Rate) - float32(player1.Rate)) / 800) + 0.5))
		player2.Rate += int(32 * (((float32(player2.Rate) - float32(player1.Rate)) / 800) + 0.5))
	}
	
	db.Model(&player1).Update("rate", player1.Rate)
	db.Model(&player2).Update("rate", player2.Rate)

	db.Delete(&token)

	fmt.Fprintf(w, "{\"error\": false}")
}