package main

import (
	"github.com/jinzhu/gorm"
)

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

func getMLBTotalPitchingStat(p Player) PitchingStat {
	for _, stat := range p.Pitching {
		if stat.Year == "MLB通算" {
			return stat
		}
	}
	return PitchingStat{}
}

func getMLBTotalBattingStat(p Player) BattingStat {
	for _, stat := range p.Batting {
		if stat.Year == "MLB通算" {
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
			if stat.MLB {
				score = score * 1.5
			}
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
			if stat.MLB {
				score = score * 1.5
			}
			if score > max {
				careerHigh = stat
				max = score
			}
		}
	}
	return careerHigh
}
