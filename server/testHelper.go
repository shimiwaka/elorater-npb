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
	player := Player{Name: name, Rate: rate}
	db.Create(&player)
	return player.ID
}
