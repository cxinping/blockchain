package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Tournament struct {
	// 赛事
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}
