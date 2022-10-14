package util

import (
	"testing"
)

func TestRandomString(t *testing.T) {

	for i := 0; i < 5; i++ {
		rst := RandomString()
		t.Log("结果是 ", rst)
	}

}
