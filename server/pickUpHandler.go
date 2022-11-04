package main

import (
	"crypto/md5"
	"net/http"
	"fmt"
	"time"
	"math/rand"
)

type pickUpHandler struct{}

func (p *pickUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int
	var player1 Player
	var player2 Player

	rand.Seed(time.Now().UnixNano())

	db := ConnectDB()
	db.Model(&Player{}).Count(&count)

	db.Limit(1).Offset(rand.Intn(count)).Find(&player1)
	max := player1.Rate + 100
	min := player1.Rate - 100

	db.Model(&Player{}).Where("rate > ?", min).Where("rate < ?", max).Count(&count)

	loop := 0
	for player2.Rate == 0 || player1.ID == player2.ID {
		db.Where("rate > ?", min).Where("rate < ?", max).Limit(1).Offset(rand.Intn(count)).Find(&player2)
		loop++
		if loop >= 100 {
			db.Limit(1).Offset(rand.Intn(count)).Find(&player2)
		}
	}

	seed := []byte(player1.Name + player2.Name + string(time.Now().UnixNano()))
	tokenString := fmt.Sprintf("%x", md5.Sum(seed))

	fmt.Fprintf(w, "%s\n", tokenString)
	fmt.Fprintf(w, "%s\n%s\n", player1.Name, player2.Name)

	token := Token{Token: tokenString, Player1_id: player1.ID, Player2_id: player2.ID}
	db.Create(&token)
}
