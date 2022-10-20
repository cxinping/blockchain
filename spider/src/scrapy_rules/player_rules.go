package scrapy_rules

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"spider/src/config"
	"spider/src/crawl"
	"spider/src/model"
	"spider/src/utils"
	"strings"
	"time"
)

func SetPlayerCallback(getPlayerC *colly.Collector, playerUrl string) {
	getPlayerC.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		player := crawl.ParseMatchTeamPlayer(dom)
		player.PlayerUrl = playerUrl
		var DB = config.GetDB() // 初始化数据库句柄
		OperatePlayer(DB, player)
	})

	getPlayerC.OnResponse(func(r *colly.Response) {
		fmt.Println("访问战队-队员网页 Visited ", r.Request.URL.String())
	})
}

func OperatePlayer(DB *gorm.DB, player model.Player) {
	// 处理战队的队员数据
	var playerCount int = 0
	var queryTeam = model.Team{}

	DB.Where("team_name = ?", player.CurrentTeamName).Find(&queryTeam)
	DB.Model(&model.Player{}).Where("nick_name = ?", player.NickName).Count(&playerCount)

	// 存在队员记录就修改，不存在就新建队员记录
	//if playerCount == 0 && queryTeam.TeamBizId != "" {
	if playerCount == 0 {
		player.PlayerBizId = utils.GenerateModuleBizID("PR")
		player.TeamBizId = queryTeam.TeamBizId
		player.CreatedTime = time.Now()
		player.Insert(DB)
	} else {
		DB.Model(model.Player{}).Where("nick_name = ?", player.NickName).Update(player)
	}

}
