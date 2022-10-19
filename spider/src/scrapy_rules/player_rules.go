package scrapy_rules

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"spider/src/config"
	"spider/src/crawl"
	"strings"
)

func SetPlayerCallback(getPlayerC *colly.Collector, playerUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

	getPlayerC.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		//fmt.Println(dom.Html())
		player := crawl.ParseMatchTeamPlayer(dom)
		fmt.Println(player)
		player.PlayerUrl = playerUrl
		crawl.OperatePlayer(DB, player)
	})

	// 异常处理
	getPlayerC.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		fmt.Println("")
	})

	getPlayerC.OnResponse(func(r *colly.Response) {
		fmt.Println("访问战队-队员网页 Visited ", r.Request.URL.String())
	})
}
