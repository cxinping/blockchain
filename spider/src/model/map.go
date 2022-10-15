package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Map struct {
	// 地图
	gorm.Model
	Map_biz_id   string
	Map_name     string
	Map_pic      string
	Created_time time.Time
}

func (Map) TableName() string {
	// 自定义表的名称
	return "map"
}

func (mapObj *Map) Insert(DB *gorm.DB) {
	DB.Table("map").Create(mapObj)
	//DB.Table("map").Debug().Create(map)
}
