package main

import (
	"bytes"
	"encoding/json"
	"crypto/md5"
	"net/http"
	"fmt"
	"time"
	"math/rand"
)

type pickUpHandler struct{}

type PickUpResponse struct {
	Token string			`json:"token"`
	Player1 PickUpPlayer	`json:"player1"`
	Player2 PickUpPlayer	`json:"player2"`
}

type PickUpPlayer struct {
	Name string						`json:"name"`
	Birth string					`json:"birth"`
	BT string						`json:"bt"`
	PitchingCareerHigh PitchingStat	`json:"pitchingCareerHigh"`
	BattingCareerHigh BattingStat	`json:"battingCareerHigh"`
	PitchingTotal PitchingStat		`json:"pitchingTotal"`
	BattingTotal BattingStat		`json:"battingTotal"`
}

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

	resp := PickUpResponse{Token: tokenString}

	resp.Player1.Name = player1.Name
	resp.Player2.Name = player2.Name
	resp.Player1.Birth = player1.Birth
	resp.Player2.Birth = player2.Birth
	resp.Player1.BT = player1.BT
	resp.Player2.BT = player2.BT

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&resp); err != nil {
		panic("encode error")
	}
	fmt.Fprint(w, buf.String())

	token := Token{Token: tokenString, Player1_id: player1.ID, Player2_id: player2.ID}
	db.Create(&token)
}
