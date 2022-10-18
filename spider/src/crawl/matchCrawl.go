package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strings"
	"time"
)

// 爬取将要进行的比赛网页数据
func CrawlMatcheWeb(matchUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

	// 爬取赛事信息
	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	//爬取赛事和赛果网页数据
	c.OnHTML("div.match-page", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

		matchTime, matchMode, matchStatus, team1, team2 := ParseMatchDetail(dom)
		//fmt.Println("matchTime=", matchTime)
		//fmt.Println("matchModeStr=", matchModeStr)
		//fmt.Println("team1.TeamUrl=", team1.TeamUrl)
		//fmt.Println(team2)
		operateMatchDetail(DB, matchUrl, matchTime, matchMode, matchStatus, team1, team2)

		CrawlTeam(team1.TeamUrl)
		CrawlTeam(team2.TeamUrl)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问比赛网页 Visited ", r.Request.URL.String())
	})

	c.Visit(matchUrl)
}

// 爬取已经结束的比赛网页数据
func CrawlMatcheResultWeb() {
	//DB := config.GetDB() // 初始化数据库句柄

	// 爬取赛事信息
	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	//爬取赛事和赛果网页数据
	c.OnHTML("div.results-holder allres", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		ParseMatchResult(dom)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问已经有比赛结果的赛果网页 Visited ", r.Request.URL.String())
	})

	c.Visit(parameter.MATCH_RESULT_URL)
}

func operateMatchDetail(DB *gorm.DB, matchUrl string, matchTime time.Time, matchMode string, matchStatus string, team1 model.Team, team2 model.Team) {
	//处理比赛详细数据
	var match = model.Match{}
	DB.Where("match_url = ?", matchUrl).Find(&match)

	// 修改比赛的状态
	if match.Status == parameter.MATCH_STATUS_LIVE {
		DB.Model(&match).Updates(model.Match{Mode: matchMode, MatchTime: matchTime, Status: matchStatus})
	} else if match.Status == parameter.MATCH_STATUS_UNSTARTED {
		DB.Model(&match).Updates(model.Match{Mode: matchMode, Status: matchStatus})
	}

	// 处理战队数据
	var count int = 0
	DB.Model(&model.Team{}).Where("team_name = ?", team1.TeamName).Count(&count)
	if count == 0 {
		team1.TeamBizId = utils.GenerateModuleBizID("TM")
		team1.CreatedTime = time.Now()
		team1.Insert(DB)
	}
	count = 0
	DB.Model(&model.Team{}).Where("team_name = ?", team2.TeamName).Count(&count)
	if count == 0 {
		team2.TeamBizId = utils.GenerateModuleBizID("TM")
		team2.CreatedTime = time.Now()
		team2.Insert(DB)
	}

	if match.Team1BizId == "" && match.Team2BizId == "" {
		DB.Model(&match).Updates(model.Match{Team1BizId: team1.TeamBizId, Team2BizId: team2.TeamBizId})
	}

}
