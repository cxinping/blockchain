package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"strings"
)

// 爬取队友的网页数据
func CrawlPlayer(playerUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	c.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		player := ParseMatchTeamPlayer(dom)
		OperatePlayer(DB, player)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问队员网页 Visited ", r.Request.URL.String())
	})

	c.Visit(playerUrl)
}

func OperatePlayer(DB *gorm.DB, player model.Player) {

}
