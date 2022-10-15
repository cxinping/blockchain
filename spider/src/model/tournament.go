package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tournament struct {
	// 赛事
	gorm.Model
	TtBizId     string
	TtName      string
	TtStartdate time.Time
	TtEnddate   time.Time
	TtUrl       string
	Desc        string
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
