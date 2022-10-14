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
	//DB.Table("tournament").Create(user)
	//defer DB.Close()

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

func (match *Match) Insert() {
	//defer DB.Close()
	//DB.Table("match").Create(match)
	DB.Table("match").Debug().Create(match)
}

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

func (team *Team) Insert() {
	//defer DB.Close()
	//DB.Table("team").Create(match)
	DB.Table("team").Debug().Create(team)
}

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

func (player *Player) Insert() {
	//defer DB.Close()
	//DB.Table("player").Create(player)
	DB.Table("player").Debug().Create(player)
}
