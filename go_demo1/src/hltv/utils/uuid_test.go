package utils

import (
	"fmt"
	"github.com/go-basic/uuid"
	"strings"
	"testing"
)

func TestGenerateUUID1(t *testing.T) {

	for i := 0; i < 10; i++ {
		uuid := uuid.New()
		uuid = strings.Replace(uuid, "-", "", -1)
		fmt.Println(uuid, len(uuid))
	}

}
