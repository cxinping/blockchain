package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Team struct {
	// 战队
	gorm.Model
	TeamBizId        string `gorm:"size:50;not null;default:'';comment:'战队的业务id'"`
	TeamName         string `gorm:"size:50;not null;default:'';comment:'战队的名字'"`
	TeamPic          string `gorm:"size:255;not null;default:'';comment:'战队的图片'"`
	TeamUrl          string `gorm:"size:255;default:'';comment:'战队的url'"`
	NationName       string
	NationPic        string
	WorldRanking     uint16 `gorm:"size:11;not null;default:0;comment:'国际排名'"`
	AveragePlayerAge float32
	CoatchBizId      string
	CoatchName       string `gorm:"size:100;default:'';comment:'战队教练'"`
	Players          []Player
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
