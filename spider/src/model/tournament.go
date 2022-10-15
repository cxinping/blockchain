package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tournament struct {
	// 赛事
	gorm.Model
	TtBizId     string `gorm:"size:50;not null;default:'';comment:'赛事的业务id'"`
	TtName      string `gorm:"size:100;not null;default:'';comment:'赛事的名称'"`
	TtStartdate time.Time
	TtEnddate   time.Time
	TtUrl       string `gorm:"size:100;not null;default:'';comment:'赛事的链接'"`
	Desc        string `gorm:"size:100;not null;default:'';comment:'描述'"`
	CreatedTime time.Time
}

func (Tournament) TableName() string {
	// 自定义表的名称
	return "tournament"
}

func (tt *Tournament) Insert(DB *gorm.DB) {
	DB.Table("tournament").Create(tt)
	//DB.Table("tournament").Debug().Create(tt)
}
