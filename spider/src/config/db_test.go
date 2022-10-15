package config

import (
	"fmt"
	"spider/src/model"
	"testing"
	"time"
)

func init() {
	InitDB()
}

func TestInitDBTables(t *testing.T) {
	//初始化表结构
	DB := GetDB()
	fmt.Println("db=> ", DB)
}

func TestAddTournament(t *testing.T) {
	DB := GetDB()
	tt := model.Tournament{
		TT_biz_id:    "aaaaa",
		TT_name:      "2222",
		Desc:         "qqqqqq",
		TT_startdate: time.Now(),
		TT_enddate:   time.Now(),
		Created_time: time.Now()}
	fmt.Println(tt)
	tt.Insert(DB)
}
