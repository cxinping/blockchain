package crawl

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

	config.InitDB()
}

func TestCrawlTournamentWeb(t *testing.T) {
	t.Log("*** 开始解析赛事和赛程网页 ***")
	CrawlTournamentWeb()
}

func TestCrawlMatcheWeb(t *testing.T) {
	t.Log("** 开始解析比赛的网页 **")
	start := time.Now()
	matchUrl := "https://www.hltv.org/matches/2359687/sc-vs-prospects-cct-north-europe-series-1"
	CrawlMatcheWeb(matchUrl)

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}

func TestCrawlTeam(t *testing.T) {
	t.Log("* 开始解析战队的网页 *")
	teamUrl := "https://www.hltv.org/team/11717/arena"
	CrawlTeam(teamUrl)

	//CrawlTeamHttp(teamUrl)   返回的网页内容没有经过JS渲染，没有找到数据???

}

func TestCrawlPlayer(t *testing.T) {
	t.Log("* 开始解析战队-队员的网页 *")
	playerUrl := "https://www.hltv.org/player/20582/kiyo"
	CrawlPlayer(playerUrl)
}

func TestCrawlMatcheResultWeb(t *testing.T) {
	t.Log("*** 开始解析已经有比赛结果的赛果网页 111")
	start := time.Now()
	//CrawlMatcheResults()

	// https://www.hltv.org/results?offset=100
	// https://www.hltv.org/results
	CrawlMatcheResultWeb("https://www.hltv.org/results")

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}
