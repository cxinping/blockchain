package test

import (
	"fmt"
	"runtime"
	"spider/src/config"
	"spider/src/scrapy_rules"
	"testing"
	"time"
)

func init() {
	fmt.Printf("本台电脑是 %d 核的CPU\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	config.InitDB()
}

func TestScrapyPlayerInfomation(t *testing.T) {
	start := time.Now()
	playerUrls := make([]string, 0)
	playerUrls = append(playerUrls, "https://www.hltv.org/player/11205/stadodo")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20463/ddias")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20465/arrozdoce")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20743/suka")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/21014/ag1l")

	for _, playerUrl := range playerUrls {
		scrapy_rules.ScrapyPlayerInfomation(playerUrl)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestScrapyTeamInformation(t *testing.T) {
	start := time.Now()
	// 单条战队页面抓取
	//teamUrl := "https://www.hltv.org/team/7379/ftw"
	//scrapy_rules.ScrapyTeamInformation(teamUrl)

	// 多条战队页面抓取
	teamUrls := make([]string, 0)
	teamUrls = append(teamUrls, "https://www.hltv.org/team/11826/vendetta")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/9943/atk")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/11948/nouns")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/7379/ftw")

	for _, teamUrl := range teamUrls {
		scrapy_rules.ScrapyTeamInformation(teamUrl)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestScrapyMatchInformation(t *testing.T) {

}
