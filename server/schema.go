package main

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model
	Name string
	Birth string
	BT string
	Rate int
	Pitching []PitchingStat `gorm:"foreignkey:PlayerID"`
	Batting []BattingStat `gorm:"foreignkey:PlayerID"`
}

type PitchingStat struct {
	gorm.Model
	Year string
	Game int
	Starter int
	CG int
	Shutout int
	Win int
	Lose int
	Save int
	Hold int
	Inning float64
	K int
	Era float64
	WHIP float64
	MLB bool
	PlayerID int
}

type BattingStat struct {
	gorm.Model
	Year string
	Game int
	PA int
	Atbat int
	Hit int
	TwoBase int
	ThreeBase int
	HR int
	RBI int
	SB int
	CS int
	SH int
	SF int
	BB int
	IW int
	HBP int
	SO int
	Avg float64
	Obp float64
	Slg float64
	Ops float64
	MLB bool
	PlayerID int
}
