package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"spider/src/model"
	"spider/src/util"
	"strings"
	"time"
)

func CrawlMatches() {
	// 爬取赛事信息
	base_url := util.MATCH_URL // "https://www.hltv.org/matches"
	//fmt.Println("*** 开始爬取hltv的赛事列表 ", base_url)

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", util.RandomString())
		//fmt.Println("OnRequest")
		//fmt.Println("url => ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))

		matchResultSet := OperateLivingMatch(dom)
		saveMatches(matchResultSet)

		//OperateUpcomingMatch(dom)
	})

	c.Visit(base_url)
}

func saveMatches(matches []model.Match) {
	// 批量保存Match
	if matches != nil {
		for _, match := range matches {
			match.Match_biz_id = util.GenerateModuleBizID("MH")
			match.Match_time = time.Now()
			match.Created_time = time.Now()
			match.Status = util.MATCH_STATUS_LIVE
			//fmt.Println(idx, match)
			//match.Insert()
		}
	}
}
