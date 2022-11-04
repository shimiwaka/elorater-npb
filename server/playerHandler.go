package main

import (
	"net/http"
	"fmt"
    "strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type playerHandler struct{}

func getCareerPitchingStat(p Player) PitchingStat {
	for _, stat := range p.Pitching {
		if stat.Year == "国内通算" {
			return stat
		}
	}
	return PitchingStat{}
}

func getCareerBattingStat(p Player) BattingStat {
	for _, stat := range p.Batting {
		if stat.Year == "国内通算" {
			return stat
		}
	}
	return BattingStat{}
}

func getPlayerAllStats(db *gorm.DB, id int) (Player, error) {
	var player Player
	err := db.Model(&Player{}).Preload("Pitching").Preload("Batting").First(&player, id).Error
	return player, err
}

func showCareer(p Player) string {
	pStat := getCareerPitchingStat(p)
	bStat := getCareerBattingStat(p)
	ret := fmt.Sprintf("%s %s\n%d勝%d敗 %.2f\n%.3f %d本 %d打点",
						p.Name, p.Birth,
						pStat.Win, pStat.Lose, pStat.Era,
						bStat.Avg, bStat.HR, bStat.RBI)

	return ret
}

func (p *playerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Show Player %s\n", vars["num"])

	id, _ := strconv.Atoi(vars["num"])
	db := ConnectDB()

	var player Player

	player, err := getPlayerAllStats(db, id)

	if err != nil {
		panic("failed to get player")
	}

	fmt.Fprintf(w, "%s", showCareer(player))

}
