package utils

import (
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
	MsToTime("1665844800000")

}
