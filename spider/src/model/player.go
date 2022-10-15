package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Player struct {
	//队员
	gorm.Model
	Player_biz_id string
	Name          string
	Birthday      string
	Total_award   int32
	Player_pic    string
	Player_age    uint8
	Nation_name   string
	Nation_pic    string
	Current_team  string
	Rating2       string
	Dpr           string
	Kast          string
	Impact        string
	Adr           string
	Kpr           string
	Job_status    string
	Created_time  time.Time
}

func (player *Player) Insert(DB *gorm.DB) {
	DB.Table("player").Create(player)
	//DB.Table("player").Debug().Create(player)
}
