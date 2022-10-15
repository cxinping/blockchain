package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Team struct {
	// 战队
	gorm.Model
	TeamBizId        string
	TeamName         string
	TeamPic          string
	NationName       string
	NationPic        string
	WorldRanking     uint16
	AveragePlayerAge float32
	CoatchBizId      string
	CreatedTime      time.Time
}

func (Team) TableName() string {
	// 自定义表的名称
	return "team"
}

func (team *Team) Insert(DB *gorm.DB) {
	DB.Table("team").Create(team)
	//DB.Table("team").Debug().Create(team)
}
