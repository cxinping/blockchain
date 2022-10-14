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
		TT_biz_id:    "aaaaa",
		TT_name:      "bbbb",
		TT_startdate: time.Now(),
		TT_enddate:   time.Now(),
		Created_time: time.Now()}
	fmt.Println(tt)
	//SaveTournament(&tt)
	tt.Insert()
}
