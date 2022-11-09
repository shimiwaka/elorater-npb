package main

import (
	"bytes"
	"encoding/json"
	"crypto/md5"
	"net/http"
	"fmt"
	"time"
	"math/rand"

	"github.com/jinzhu/gorm"
)

type pickUpHandler struct{}

type PickUpResponse struct {
	Error bool				`json:"error"`
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
	PitchingMLBTotal PitchingStat	`json:"pitchingMLBTotal"`
	BattingMLBTotal BattingStat		`json:"battingMLBTotal"`
}

func pickUp(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var count1, count2 int
	var player1, player2 Player

	rand.Seed(time.Now().UnixNano())
	db.Model(&Player{}).Count(&count1)

	if count1 <= 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"no player data in database\"}")
		return
	}

	db.Limit(1).Offset(rand.Intn(count1)).Find(&player1)
	max := player1.Rate + 50
	min := player1.Rate - 50
	db.Model(&Player{}).Where("rate > ?", min).Where("rate < ?", max).Where("id != ?", player1.ID).Count(&count2)

	if count2 <= 1 {
		max = 9999
		min = -9999
		count2 = count1 - 1
	}

	loop := 0
	for player2.Rate == 0  {
		db.Where("rate > ?", min).Where("rate < ?", max).Where("id != ?", player1.ID).Limit(1).Offset(rand.Intn(count2)).Find(&player2)
		loop++
		if loop >= 1000 {
			player2 = Player{}
			break
		}
	}

	player1, err := getPlayerAllStats(db, player1.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to retrieve player infomation\"}")
		return
	}

	player2, err = getPlayerAllStats(db, player2.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to retrieve player infomation\"}")
		return
	}

	seed := []byte(player1.Name + player2.Name + fmt.Sprint(time.Now().UnixNano()))
	tokenString := fmt.Sprintf("%x", md5.Sum(seed))

	resp := PickUpResponse{Error: false, Token: tokenString}

	resp.Player1.Name = player1.Name
	resp.Player2.Name = player2.Name
	resp.Player1.Birth = player1.Birth
	resp.Player2.Birth = player2.Birth
	resp.Player1.BT = player1.BT
	resp.Player2.BT = player2.BT

	resp.Player1.PitchingTotal = getTotalPitchingStat(player1)
	resp.Player2.PitchingTotal = getTotalPitchingStat(player2)
	resp.Player1.BattingTotal = getTotalBattingStat(player1)
	resp.Player2.BattingTotal = getTotalBattingStat(player2)

	resp.Player1.PitchingMLBTotal = getMLBTotalPitchingStat(player1)
	resp.Player2.PitchingMLBTotal = getMLBTotalPitchingStat(player2)
	resp.Player1.BattingMLBTotal = getMLBTotalBattingStat(player1)
	resp.Player2.BattingMLBTotal = getMLBTotalBattingStat(player2)

	resp.Player1.BattingCareerHigh = getCareerHighBattingStat(player1)
	resp.Player2.BattingCareerHigh = getCareerHighBattingStat(player2)
	resp.Player1.PitchingCareerHigh = getCareerHighPitchingStat(player1)
	resp.Player2.PitchingCareerHigh = getCareerHighPitchingStat(player2)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err = enc.Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to encode json\"}")
		return
	}
	fmt.Fprint(w, buf.String())

	token := Token{Token: tokenString, Player1_id: player1.ID, Player2_id: player2.ID}
	db.Create(&token)
}

func (p *pickUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	pickUp(db, w, r)
}
