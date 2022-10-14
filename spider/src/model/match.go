package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Match struct {
	//赛果/赛程
	gorm.Model
	Match_biz_id        string
	Match_url           string
	TT_pic              string
	TT_biz_id           string
	TT_name             string
	Status              string
	Result              string
	Mode                string
	Match_time          time.Time
	Team1_biz_id        string
	Team2_biz_id        string
	Team1_playing_score uint16
	Team2_playing_score uint16
	Team1_win_score     uint16
	Team2_win_score     uint16
	Map_type            string
	Suggest_idx         uint8
	Created_time        time.Time
	Desc                string
}

func (match *Match) Insert(DB *gorm.DB) {
	//defer DB.Close()
	DB.Table("match").Create(match)
	//DB.Table("match").Debug().Create(match)
}
