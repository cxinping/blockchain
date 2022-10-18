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
	matchUrl := "https://www.hltv.org/matches/2359684/havu-vs-ikla-cct-north-europe-series-1"
	CrawlMatcheWeb(matchUrl)
}

func TestCrawlTeam(t *testing.T) {
	t.Log("*** 开始解析战队的网页 ***")
	// https://www.hltv.org/team/7865/havu
	// https://www.hltv.org/team/11982/ikla
	teamUrl := "https://www.hltv.org/team/11982/ikla"
	CrawlTeam(teamUrl)
}

func TestCrawlPlayer(t *testing.T) {
	t.Log("*** 开始解析战队-队员的网页 ***")
	playerUrl := "https://www.hltv.org/player/18227/sensei"
	CrawlPlayer(playerUrl)
}
