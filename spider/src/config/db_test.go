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
		TtBizId:     "aaaaa",
		TtName:      "2222",
		Desc:        "qqqqqq",
		TtStartdate: time.Now(),
		TtEnddate:   time.Now(),
		CreatedTime: time.Now()}
	fmt.Println(tt)
	tt.Insert(DB)
}
