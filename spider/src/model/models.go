package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tournament struct {
	// 赛事
	gorm.Model
	TT_biz_id    string
	TT_name      string
	TT_startdate time.Time
	TT_enddate   time.Time
	Desc         string
	Created_time time.Time
}

func (Tournament) TableName() string {
	// 自定义表的名称
	return "tournament"
}

func (tt *Tournament) Insert() {
	//db.Table("user").Create(user)
	DB.Table("tournament").Debug().Create(tt)
}

type Match struct {
	//赛果/赛程
	gorm.Model
	Match_biz_id        string
	Match_url           string
	TT_pic              string
	TT_biz_id           string
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
}

func (Match) TableName() string {
	// 自定义表的名称
	return "match"
}

type Team struct {
	// 战队
	//赛果/赛程
	gorm.Model
	Team_biz_id string
	Team_name   string
	Team_pic    string
	Nation_name string
	Nation_pic  string
}

func (Team) TableName() string {
	//自定义表的名称
	return "team"
}
