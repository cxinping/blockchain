package spider

import (
	"fmt"
	"github.com/gocolly/colly"
)

func CrawlMatches() {
	// 爬取赛事信息
	base_url := "https://www.hltv.org/matches"
	fmt.Println("*** 开始爬取hltv的赛事列表 ", base_url)

	c := colly.NewCollector()

	// selector goquery name id class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))

		ret, _ := e.DOM.Html()
		fmt.Println("ret-> ", ret)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url => ", r.URL)
	})

	c.Visit("base_url")

}
