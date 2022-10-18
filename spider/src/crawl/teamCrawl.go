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
	"time"
)

// 爬取战队的网页数据
func CrawlTeam(teamUrl string) {
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
		team := ParseMatchTeam(dom)
		team.TeamUrl = teamUrl
		operateMatchTeam(DB, team)

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited ", r.Request.URL.String())
	})

	c.Visit(teamUrl)
}

func operateMatchTeam(DB *gorm.DB, team model.Team) {
	// 处理比赛战队
	//fmt.Println(team)

	var count int = 0
	DB.Model(&model.Team{}).Where("team_name = ?", "abc").Count(&count)
	if count == 0 {
		fmt.Println("111 ", team.AveragePlayerAge)
		team.TeamBizId = utils.GenerateModuleBizID("TM")
		team.CreatedTime = time.Now()
		team.Insert(DB)
	} else {
		fmt.Println("222")
	}
}
