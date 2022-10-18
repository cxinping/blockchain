package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Match struct {
	//赛果/赛程
	gorm.Model
	MatchBizId        string    `gorm:"size:50;not null;default:'';comment:'赛程的名称'"`
	MatchUrl          string    `gorm:"size:255;default:'';comment:'赛程的链接'"`
	TtPic             string    `gorm:"size:255;default:'';comment:'赛事的图片'"`
	TtBizId           string    `gorm:"size:50;not null;default:'';comment:'赛事的业务id'"`
	TtName            string    `gorm:"size:255;default:'';comment:'比赛的名称'"`
	Status            string    `gorm:"size:10;default:'';comment:'比赛状态'"`
	Result            string    `gorm:"size:100;default:'';comment:'比赛结果'"`
	Mode              string    `gorm:"size:50;default:'';comment:'比赛模式: 线上/线下'"`
	MatchTime         time.Time `gorm:"default:null;comment:'比赛时间'"`
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
	Desc              string `gorm:"size:255;default:'';comment:'描述'"`
	Team1             Team
	Team2             Team
}

func (Match) TableName() string {
	// 自定义表的名称
	return "match"
}

func (match *Match) Insert(DB *gorm.DB) {
	DB.Table("match").Create(match)
	//DB.Table("match").Debug().Create(match)
}
