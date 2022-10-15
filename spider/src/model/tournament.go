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
	TT_url       string
	Desc         string
	Created_time time.Time
}

func (Tournament) TableName() string {
	// 自定义表的名称
	return "tournament"
}

func (tt *Tournament) Insert(DB *gorm.DB) {
	DB.Table("tournament").Create(tt)
	//DB.Table("tournament").Debug().Create(tt)
}
