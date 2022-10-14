package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	DB = GetDBInstance()
}

func InitTables() {
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

	//赛事
	db.AutoMigrate(&Tournament{})
	//赛程/赛果
	db.AutoMigrate(&Match{})
}

func GetDBInstance() *gorm.DB {
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
	return db
}

func SaveTournament(tt *Tournament) {
	DB.Table("tournament").Debug().Create(tt)
}
