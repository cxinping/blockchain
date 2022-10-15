package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Player struct {
	//队员
	gorm.Model
	PlayerBizId string `gorm:"size:50;not null;default:'';comment:'队员的业务id'"`
	Name        string `gorm:"size:50;not null;default:'';comment:'队员的姓名'"`
	Birthday    string `gorm:"size:50;default:'';comment:'队员的生日'"`
	TotalAward  int32  `gorm:"size:10;default:0;comment:'总奖金'"`
	PlayerPic   string
	PlayerAge   uint8  `gorm:"size:10;default:0;comment:'年龄'"`
	NationName  string `gorm:"size:50;default:'';comment:'队员的国籍'"`
	NationPic   string `gorm:"size:100;default:'';comment:'队员的国籍图片'"`
	CurrentTeam string `gorm:"size:50;default:'';comment:'队员所属的战队'"`
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
