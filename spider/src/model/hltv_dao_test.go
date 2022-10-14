package model

import (
	"fmt"
	"testing"
	"time"
)

func TestInitTable(t *testing.T) {
	InitTables()
}

func TestAddTournament(t *testing.T) {
	tt := Tournament{
		TT_biz_id:    "123456",
		TT_name:      "ccccc",
		TT_startdate: time.Now(),
		TT_enddate:   time.Now(),
		Created_time: time.Now()}
	fmt.Println(tt)
	//SaveTournament(&tt)
	tt.Insert()
}
