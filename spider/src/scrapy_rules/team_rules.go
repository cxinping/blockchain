package scrapy_rules

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strconv"
	"strings"
	"time"
)

func SetTeamCallback(getTeamC *colly.Collector, teamUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

	getTeamC.OnResponse(func(r *colly.Response) {
		idx := strings.Index(r.Request.URL.String(), "team")
		if idx > -1 {
			fmt.Println("访问战队网页 Visited ", r.Request.URL.String())
		} else {
			fmt.Println("访问战队-队员网页 Visited ", r.Request.URL.String())
		}

		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))
		team := ParseMatchTeam(dom)
		team.TeamUrl = teamUrl
		OperateMatchTeam(DB, team)

		if len(team.Players) > 0 {
			for _, player := range team.Players {
				getTeamC.Visit(player.PlayerUrl)
			}
		}
	})

	getTeamC.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		requestUrl := e.Request.URL.String()
		idx := strings.Index(requestUrl, "player")
		//fmt.Println("SetTeamCallback2 OnHTML requestUrl=> ", requestUrl, ", idx=", idx)

		if idx > -1 {
			content, _ := e.DOM.Html()
			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
			player := ParseMatchTeamPlayer(dom)
			player.PlayerUrl = requestUrl
			OperatePlayer(DB, player)
		}
	})

}

func ParseMatchTeam(dom *goquery.Document) model.Team {
	//解析比赛战队网页
	var team model.Team
	playerResultSet := make([]model.Player, 0)

	teamProfileDom := dom.Find("div[class='profile-team-stats-container']").Find("div[class='profile-team-stat']")
	worldRankingStr := teamProfileDom.Eq(0).Find("span").Text()
	worldRanking, _ := strconv.Atoi(strings.Replace(worldRankingStr, "#", "", -1))

	averagePlayerAge := float64(0)
	coatchName := ""
	itemStr := utils.CompressString(teamProfileDom.Eq(2).Find("b").Text())
	//fmt.Println("**** itemStr=", itemStr)
	if itemStr == "Coach" {
		coatchName = teamProfileDom.Eq(2).Find("span").Text()
		coatchName = strings.Replace(coatchName, "'", "", -1)
		//fmt.Println("111 coatchName=", coatchName)
	} else {
		averagePlayerAgeStr := teamProfileDom.Eq(2).Find("span").Text()
		averagePlayerAge, _ = strconv.ParseFloat(averagePlayerAgeStr, 64)
		averagePlayerAge = utils.Decimal(averagePlayerAge)
		coatchName = teamProfileDom.Eq(3).Find("span").Text()
		coatchName = strings.Replace(coatchName, "'", "", -1)
	}

	//fmt.Println("222 coatchName=", coatchName, ", worldRanking=", worldRanking)

	teamName := dom.Find("div[class='profile-team-info']").Find("h1[class='profile-team-name text-ellipsis']").Text()
	profileTopDom := dom.Find("div[class='standard-box profileTopBox clearfix']").Find("div[class='flex']")
	teamPic, _ := profileTopDom.Find("div[class='profile-team-logo-container']").Find("img").Attr("src")
	nationName := profileTopDom.Find("div[class='profile-team-info']").Find("div[class='team-country text-ellipsis']").Text()
	nationPic, _ := profileTopDom.Find("div[class='profile-team-info']").Find("div[class='team-country text-ellipsis']").Find("img").Attr("src")
	nationPic = parameter.HLTV_INDEX + nationPic

	team.TeamName = utils.CompressString(teamName)
	team.WorldRanking = uint16(worldRanking)
	team.AveragePlayerAge = float32(averagePlayerAge)
	team.CoatchName = utils.CompressString(coatchName)
	team.TeamPic = parameter.HLTV_INDEX + teamPic
	team.NationName = utils.CompressString(nationName)
	team.NationPic = nationPic

	dom.Find("div[class='bodyshot-team g-grid']").Find("a[class='col-custom']").Each(func(idx int, selection *goquery.Selection) {
		var player model.Player
		playerUrl, _ := selection.Attr("href")
		playerUrl = parameter.HLTV_INDEX + playerUrl
		player.PlayerUrl = playerUrl
		playerPic, _ := selection.Find("img").Attr("src")
		playerName := selection.Find("div[class='playerFlagName']").Find("span[class='text-ellipsis bold']").Text()
		player.PlayerPic = playerPic
		player.NickName = utils.CompressString(playerName)
		nationPic, _ := selection.Find("span[class='gtSmartphone-only']").Find("img").Attr("src")
		nationPic = parameter.HLTV_INDEX + nationPic
		player.NationPic = nationPic
		player.CurrentTeamName = utils.CompressString(teamName)
		player.CurrentTeamPic = teamPic
		//fmt.Println(idx, playerName, nationPic)

		if playerUrl != "" {
			playerResultSet = append(playerResultSet, player)
		}

	})
	team.Players = playerResultSet

	return team
}

func OperateMatchTeam(DB *gorm.DB, team model.Team) {
	// 处理比赛战队
	var count int = 0
	DB.Model(&model.Team{}).Where("team_url = ?", team.TeamUrl).Count(&count)
	// 存在战队记录就修改，不存在就新建战队记录
	if count == 0 {
		team.TeamBizId = utils.GenerateModuleBizID("TM")
		team.CreatedTime = time.Now()
		team.Insert(DB)
	} else {
		DB.Model(model.Team{}).Where("team_url = ?", team.TeamUrl).Updates(team)
	}

	// 处理战队相关的队员
	var queryTeam = model.Team{}
	DB.Where("team_url = ?", team.TeamUrl).Find(&queryTeam)
	//fmt.Println(queryTeam)
	//fmt.Println(queryTeam.TeamBizId)

	if len(team.Players) > 0 {
		for _, player := range team.Players {
			count = 0
			DB.Model(&model.Player{}).Where("player_url = ?", player.PlayerUrl).Count(&count)
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
