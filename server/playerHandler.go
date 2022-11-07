package main

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"bytes"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type playerHandler struct{}

func getTotalPitchingStat(p Player) PitchingStat {
	for _, stat := range p.Pitching {
		if stat.Year == "国内通算" {
			return stat
		}
	}
	return PitchingStat{}
}

func getTotalBattingStat(p Player) BattingStat {
	for _, stat := range p.Batting {
		if stat.Year == "国内通算" {
			return stat
		}
	}
	return BattingStat{}
}

func getPlayerAllStats(db *gorm.DB, id uint) (Player, error) {
	var player Player
	err := db.Model(&Player{}).Preload("Pitching").Preload("Batting").First(&player, id).Error
	return player, err
}

func getCareerHighBattingStat(p Player) BattingStat {
	var score float64
	var max float64
	max = -9999

	careerHigh := BattingStat{}
	for _, stat := range p.Batting {
		if stat.Year != "国内通算" && stat.Year != "MLB通算" {
			score = (stat.Ops - 0.500) * float64(stat.PA)
			if score > max {
				careerHigh = stat
				max = score
			}
		}
	}
	return careerHigh
}

func getCareerHighPitchingStat(p Player) PitchingStat {
	var score float64
	var max float64
	max = -9999

	careerHigh := PitchingStat{}
	for _, stat := range p.Pitching {
		if stat.Year != "国内通算" && stat.Year != "MLB通算" {
			score = (5 - stat.Era) * float64(stat.Inning)
			if score > max {
				careerHigh = stat
				max = score
			}
		}
	}
	return careerHigh
}

func (p *playerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["num"])
	db := ConnectDB()

	var player Player

	player, err := getPlayerAllStats(db, uint(id))

	if err != nil {
		panic("failed to get player")
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&player); err != nil {
		panic("encode error")
	}
	fmt.Fprint(w, buf.String())
}
