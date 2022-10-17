package crawl

import (
	"spider/src/config"
	"testing"
)

func init() {
	config.InitDB()
}

func TestCrawlTournamentWeb(t *testing.T) {
	t.Log("*** 开始解析赛事网页 ***")
	CrawlTournamentWeb()
}

func TestCrawlMatcheWeb(t *testing.T) {
	t.Log("*** 开始解析比赛的网页 ***")
	matchUrl := "https://www.hltv.org/matches/2359479/flames-ascent-vs-furia-academy-weplay-academy-league-season-6"
	CrawlMatcheWeb(matchUrl)

}
