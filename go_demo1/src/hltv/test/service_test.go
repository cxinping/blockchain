package test

import (
	"go_demo1/src/hltv/service"
	"testing"
)

func TestAdd(t *testing.T) {
	rst := service.Add(1, 2)
	t.Log("结果是 ", rst)
}
