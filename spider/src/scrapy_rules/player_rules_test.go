package scrapy_rules

import (
	"fmt"
	"runtime"
	"spider/src/config"
	"testing"
)

func init() {
	fmt.Printf("本台电脑是 %d 核的CPU\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	config.InitDB()
}

func TestSetPlayerCallback(t *testing.T) {
	getPlayerC := GetDefaultCollector()
	playerUrl := "https://www.hltv.org/player/8565/hen1"
	SetPlayerCallback(getPlayerC, playerUrl)

	getPlayerC.Visit(playerUrl)

	getPlayerC.Wait()
}
