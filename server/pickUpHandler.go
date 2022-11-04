package main

import (
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
	max := player1.Rate + 50
	min := player1.Rate - 50

	db.Model(&Player{}).Where("rate > ?", min).Where("rate < ?", max).Count(&count)

	db.Where("rate > ?", min).Where("rate < ?", max).Limit(1).Offset(rand.Intn(count)).Find(&player2)

	fmt.Fprintf(w, "%s\n%s\n", player1.Name, player2.Name)
}
