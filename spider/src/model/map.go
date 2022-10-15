package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Map struct {
	// 地图
	gorm.Model
	MapBizId    string `gorm:"size:50;not null;default:'';comment:'地图的业务id'"`
	MapName     string `gorm:"size:100;not null;default:'';comment:'地图的名称'"`
	MapPic      string `gorm:"size:100;not null;default:'';comment:'地图的图片'"`
	CreatedTime time.Time
}

func (Map) TableName() string {
	// 自定义表的名称
	return "map"
}

func (mapObj *Map) Insert(DB *gorm.DB) {
	DB.Table("map").Create(mapObj)
	//DB.Table("map").Debug().Create(map)
}
