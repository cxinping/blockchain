package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"spider/src/model"
	"spider/src/utils"
	"strings"
)

// 爬取战队的网页数据
func CrawlTeam(matchUrl string) {
	//DB := config.GetDB() // 初始化数据库句柄

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	c.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		team := ParseMatchTeam(dom)
		operateMatchTeam(team)

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited ", r.Request.URL.String())
	})

	c.Visit(matchUrl)
}

func operateMatchTeam(team model.Team) {
	// 处理比赛战队

}
