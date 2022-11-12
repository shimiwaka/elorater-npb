package main

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model `json:"-"`
	Name       string         `json:"name"`
	Birth      string         `json:"birth"`
	BT         string         `json:"bt"`
	Rate       int            `json:"rate"`
	Pitching   []PitchingStat `json:"pitching"`
	Batting    []BattingStat  `json:"batting"`
}

type PitchingStat struct {
	gorm.Model `json:"-"`
	Year       string  `json:"year"`
	Game       int     `json:"game"`
	Starter    int     `json:"starter"`
	CG         int     `json:"cg"`
	Shutout    int     `json:"shutout"`
	Win        int     `json:"win"`
	Lose       int     `json:"lose"`
	Save       int     `json:"save"`
	Hold       int     `json:"hold"`
	Inning     float64 `json:"inning"`
	K          int     `json:"k"`
	Era        float64 `json:"era"`
	WHIP       float64 `json:"whip"`
	MLB        bool    `json:"mlb"`
	PlayerID   int     `json:"-"`
}

type BattingStat struct {
	gorm.Model `json:"-"`
	Year       string  `json:"year"`
	Game       int     `json:"game"`
	PA         int     `json:"pa"`
	Atbat      int     `json:"atbat"`
	Hit        int     `json:"hit"`
	TwoBase    int     `json:"twoBase"`
	ThreeBase  int     `json:"threeBase"`
	HR         int     `json:"hr"`
	RBI        int     `json:"rbi"`
	SB         int     `json:"sb"`
	CS         int     `json:"cs"`
	SH         int     `json:"sh"`
	SF         int     `json:"sf"`
	BB         int     `json:"bb"`
	IW         int     `json:"iw"`
	HBP        int     `json:"hbp"`
	SO         int     `json:"so"`
	Avg        float64 `json:"avg"`
	Obp        float64 `json:"obp"`
	Slg        float64 `json:"slg"`
	Ops        float64 `json:"ops"`
	MLB        bool    `json:"mlb"`
	PlayerID   int     `json:"-"`
}

type Settings struct {
	DB_username string
	DB_pass     string
	DB_host     string
	DB_port     int
	DB_name     string
}

type Token struct {
	gorm.Model
	Player1_id uint
	Player2_id uint
	Token      string
}
