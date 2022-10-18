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
	var count int = 0
	DB.Model(&model.Team{}).Where("team_name = ?", team.TeamName).Count(&count)
	// 存在战队记录就修改，不存在就新建战队记录
	if count == 0 {
		team.TeamBizId = utils.GenerateModuleBizID("TM")
		team.CreatedTime = time.Now()
		team.Insert(DB)
	} else {
		DB.Model(model.Team{}).Where("team_name = ?", team.TeamName).Updates(team)
	}

	// 处理战队相关的队员
	var queryTeam = model.Team{}
	DB.Where("team_name = ?", team.TeamName).Find(&queryTeam)
	//fmt.Println(queryTeam)
	//fmt.Println(queryTeam.TeamBizId)

	if len(team.Players) > 0 {
		for idx, player := range team.Players {
			count = 0
			DB.Model(&model.Player{}).Where("name = ?", player.Name).Count(&count)
			fmt.Println(idx, player.Name, count)

		}
	}

}
