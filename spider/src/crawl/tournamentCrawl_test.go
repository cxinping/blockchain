package crawl

import (
	"spider/src/config"
	"testing"
)

func init() {
	config.InitDB()
}

func TestCrawlTournamentWeb(t *testing.T) {
	t.Log("*** 开始解析赛事和赛程网页 ***")
	CrawlTournamentWeb()
}

func TestCrawlMatcheWeb(t *testing.T) {
	t.Log("** 开始解析比赛的网页 **")
	matchUrl := "https://www.hltv.org/matches/2359362/dynasty-vs-arena-esl-australia-nz-championship-season-15"
	CrawlMatcheWeb(matchUrl)
}

func TestCrawlTeam(t *testing.T) {
	t.Log("* 开始解析战队的网页 *")
	teamUrl := "https://www.hltv.org/team/11717/arena"
	CrawlTeam(teamUrl)
}

func TestCrawlPlayer(t *testing.T) {
	t.Log("* 开始解析战队-队员的网页 *")
	playerUrl := "https://www.hltv.org/player/20582/kiyo"
	CrawlPlayer(playerUrl)
}

func TestCrawlMatcheResultWeb(t *testing.T) {
	t.Log("*** 开始解析已经有比赛结果的赛果网页 ***")
	CrawlMatcheResultWeb()
}
