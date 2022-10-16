package utils

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 5; i++ {
		rst := RandomString()
		t.Log("结果是 ", rst)
	}
}

func TestGenerateModuleBizID(t *testing.T) {
	for i := 0; i < 15; i++ {
		biz_id := GenerateModuleBizID("MH")
		t.Log("模块的业务ID是 ", biz_id)
	}
}

func TestMsToTime(t *testing.T) {
	tm, _ := MsToTime("1665844800000")
	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))
}
