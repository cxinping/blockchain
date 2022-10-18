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
		OperateMatchTeam(DB, team)

		if len(team.Players) > 0 {
			for _, player := range team.Players {
				//fmt.Println("player.PlayerUrl=> ", player.PlayerUrl)
				CrawlPlayer(player.PlayerUrl)
			}
		}

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问战队网页 Visited ", r.Request.URL.String())
	})

	c.Visit(teamUrl)
}

func OperateMatchTeam(DB *gorm.DB, team model.Team) {
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
		for _, player := range team.Players {
			count = 0
			DB.Model(&model.Player{}).Where("nick_name = ?", player.NickName).Count(&count)
			//fmt.Println(player.Name, count)

			if count == 0 {
				player.TeamBizId = team.TeamBizId
				player.PlayerBizId = utils.GenerateModuleBizID("PR")
				player.CreatedTime = time.Now()
				player.Insert(DB)
			}
		}
	}

}
