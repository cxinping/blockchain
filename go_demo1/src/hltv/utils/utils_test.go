package utils

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	rst := RandomString()
	for i := 0; i < 5; i++ {
		t.Log("结果是 ", rst)
	}

}
