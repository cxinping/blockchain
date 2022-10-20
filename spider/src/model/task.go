package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ErrorInfo struct {
	// 异常信息
	gorm.Model
	Url         string `gorm:"size:100;default:'';comment:'执行任务时发生异常的链接'"`
	ErrorType   string `gorm:"size:100;default:'';comment:'异常的类型'"`
	Desc        string `gorm:"size:300;default:'';comment:'异常的描述'"`
	CreatedTime time.Time
}

func (ErrorInfo) TableName() string {
	// 自定义表的名称
	return "error_info"
}

func (errInfo *ErrorInfo) Insert(DB *gorm.DB) {
	DB.Table("error_info").Create(errInfo)
	//DB.Table("error_info").Debug().Create(errInfo)
}
