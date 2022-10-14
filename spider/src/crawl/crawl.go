package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"spider/src/util"
	"strings"
)

func CrawlMatches() {
	// 爬取赛事信息
	base_url := "https://www.hltv.org/matches"
	fmt.Println("*** 开始爬取hltv的赛事列表 ", base_url)

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", util.RandomString())
		//fmt.Println("OnRequest")
		fmt.Println("url => ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))

		operate_living_match(dom)
		//operate_upcoming_match(dom)

	})

	c.Visit(base_url)
}
