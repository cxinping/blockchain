package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Player struct {
	//队员
	gorm.Model
	PlayerBizId string
	Name        string
	Birthday    string
	TotalAward  int32
	PlayerPic   string
	PlayerAge   uint8
	NationName  string
	NationPic   string
	CurrentTeam string
	Rating2     string
	Dpr         string
	Kast        string
	Impact      string
	Adr         string
	Kpr         string
	JobStatus   string
	CreatedTime time.Time
}

func (Player) TableName() string {
	// 自定义表的名称
	return "player"
}

func (player *Player) Insert(DB *gorm.DB) {
	DB.Table("player").Create(player)
	//DB.Table("player").Debug().Create(player)
}
