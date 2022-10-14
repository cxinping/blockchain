package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Team struct {
	// 战队
	gorm.Model
	Team_biz_id        string
	Team_name          string
	Team_pic           string
	Nation_name        string
	Nation_pic         string
	World_ranking      uint16
	Average_player_age float32
	Coatch_biz_id      string
	Created_time       time.Time
}

func (team *Team) Insert(DB *gorm.DB) {
	//defer DB.Close()
	//DB.Table("team").Create(match)
	DB.Table("team").Debug().Create(team)
}
