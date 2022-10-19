package scrapy_rules

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"spider/src/config"
	"spider/src/crawl"
	"strings"
)

var DB = config.GetDB() // 初始化数据库句柄

func SetPlayerCallback(getPlayerC *colly.Collector, playerUrl string) {
	getPlayerC.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		player := crawl.ParseMatchTeamPlayer(dom)
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
