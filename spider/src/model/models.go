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
	// 自定义表明
	return "tournament"
}

func (tt *Tournament) Insert() {
	//db.Table("user").Create(user)
	DB.Table("tournament").Debug().Create(tt)
}
