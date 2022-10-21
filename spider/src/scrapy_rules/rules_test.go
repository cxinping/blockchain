package scrapy_rules

import (
	"fmt"
	"runtime"
	"spider/src/config"
	"testing"
	"time"
)

func init() {
	fmt.Printf("本台电脑是 %d 核的CPU\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 初始化数据库
	config.InitDB()
}

func TestScrapyPlayer(t *testing.T) {
	start := time.Now()
	playerUrls := make([]string, 0)
	playerUrls = append(playerUrls, "https://www.hltv.org/player/11205/stadodo")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20463/ddias")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20465/arrozdoce")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20743/suka")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/21014/ag1l")

	for _, playerUrl := range playerUrls {
		ScrapyPlayerInfomation(playerUrl)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestScrapyTeam(t *testing.T) {
	start := time.Now()
	// 单条战队页面抓取
	// 该函数执行完成耗时： 41.03365201s
	//teamUrl := "https://www.hltv.org/team/11826/vendetta"
	//ScrapyTeamInformation(teamUrl)

	// 多条战队页面抓取
	teamUrls := make([]string, 0)
	teamUrls = append(teamUrls, "https://www.hltv.org/team/5347/bluejays")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/11826/vendetta")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/9943/atk")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/11948/nouns")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/7379/ftw")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/6947/teamone")
	teamUrls = append(teamUrls, "https://www.hltv.org/team/10462/brazen")

	// 该函数执行完成耗时： 5m2.679026154s
	for _, teamUrl := range teamUrls {
		ScrapyTeamInformation(teamUrl)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestScrapyMatchInformation(t *testing.T) {
	start := time.Now()

	//单条比赛数据抓取
	//matchUrl := "https://www.hltv.org/matches/2359380/vitality-vs-tyloo-blast-premier-fall-showdown-2022-europe"
	//ScrapyMatchInformation(matchUrl)

	// 多条战队页面抓取
	matchUrls := make([]string, 0)
	matchUrls = append(matchUrls, "https://www.hltv.org/matches/2359657/ftw-vs-9ine-cct-central-europe-series-3")
	matchUrls = append(matchUrls, "https://www.hltv.org/matches/2359380/vitality-vs-tyloo-blast-premier-fall-showdown-2022-europe")
	matchUrls = append(matchUrls, "https://www.hltv.org/matches/2359713/nouns-vs-atk-esl-challenger-league-season-42-north-america")

	for _, matchUrl := range matchUrls {
		ScrapyMatchInformation(matchUrl)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestScrapyMatchResults(t *testing.T) {
	// 分页爬取比赛结果列表页面的比赛数据
	start := time.Now()

	ScrapyMatchResults()

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestScrapyMatchResult(t *testing.T) {
	// 爬取比赛结果列表页面的比赛数据，每页比赛记录有100条比赛结果
	// https://www.hltv.org/results
	start := time.Now()

	matchResultUrl := "https://www.hltv.org/results"
	ScrapyMatchResult(matchResultUrl)

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

//////////////////////////////// 赛事/赛程列表页面解析
