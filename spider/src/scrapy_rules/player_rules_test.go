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
	//playerUrl := "https://www.hltv.org/player/8565/hen1"

	playerUrls := make([]string, 0)
	playerUrls = append(playerUrls, "https://www.hltv.org/player/11205/stadodo")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20463/ddias")

	for _, playerUrl := range playerUrls {
		SetPlayerCallback(getPlayerC, playerUrl)
		getPlayerC.Visit(playerUrl)
	}

	getPlayerC.Wait()
}