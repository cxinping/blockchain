package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"spider/src/model"
)

var DB *gorm.DB

// 初始化db
func InitDB() *gorm.DB {
	databaseType := "mysql"
	username := "root"      //账号
	password := "123456"    //密码
	host := "192.168.11.12" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "hltv"        //数据库名
	timeout := "10s"        //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//使用gorm链接数据库
	db, err := gorm.Open(databaseType, dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err) //数据库链接失败是致命的错误，链接失败后可以关闭程序了，所以使用logging.Fatal方法
	}

	//设置全局表名禁用复数
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)  // 用于设置闲置的连接数
	db.DB().SetMaxOpenConns(100) // 用于设置最大打开的连接数，默认值为0表示不限制

	db.AutoMigrate(&model.Tournament{}) //赛事
	db.AutoMigrate(&model.Match{})      //赛程/赛果
	db.AutoMigrate(&model.Team{})       //团队
	db.AutoMigrate(&model.Player{})     //队员

	db.LogMode(true)
	DB = db

	return db
}

// 获取db句柄
func GetDB() *gorm.DB {
	return DB
}
