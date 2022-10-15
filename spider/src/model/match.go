package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Match struct {
	//赛果/赛程
	gorm.Model
	MatchBizId        string
	MatchUrl          string
	TtPic             string
	TtBizId           string
	TtName            string
	Status            string
	Result            string
	Mode              string
	MatchTime         time.Time
	Team1BizId        string
	Team2BizId        string
	Team1Name         string
	Team2Name         string
	Team1PlayingScore uint16
	Team2PlayingScore uint16
	Team1WinScore     uint16
	Team2WinScore     uint16
	MapType           string
	SuggestIdx        uint8
	CreatedTime       time.Time
	Desc              string
}

func (Match) TableName() string {
	// 自定义表的名称
	return "match"
}

func (match *Match) Insert(DB *gorm.DB) {
	DB.Table("match").Create(match)
	//DB.Table("match").Debug().Create(match)
}
