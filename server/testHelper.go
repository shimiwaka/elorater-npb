package main

import (
	"github.com/jinzhu/gorm"
)

type DummyPlayer struct {
	Name string
	Rate int
}

func initializeDB(db *gorm.DB) {
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&PitchingStat{})
	db.AutoMigrate(&BattingStat{})
	db.AutoMigrate(&Token{})

	db.Exec("DELETE FROM players")
	db.Exec("DELETE FROM pitching_stats")
	db.Exec("DELETE FROM batting_stats")
	db.Exec("DELETE FROM tokens")
}

func setDummyPlayer(db *gorm.DB, name string, rate int) uint {
	batting := BattingStat{Year: "2000", Avg: 0.300}
	pitching := PitchingStat{Year: "2000", Era: 2.00}
	player := Player{Name: name, Rate: rate,
					Pitching: []PitchingStat{pitching},
					Batting: []BattingStat{batting},
					}
	db.Create(&player)
	return player.ID
}

func setDummyToken(db *gorm.DB) {
	token := Token{Token: "DUMMY",
					Player1_id: setDummyPlayer(db, "dummy1", 1500),
					Player2_id: setDummyPlayer(db, "dummy2", 1500)}
	db.Create(&token)
}