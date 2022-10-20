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

func SetTeamCallback(getTeamC *colly.Collector, teamUrl string, scrapyPlayer func(playerUrl string)) {
	DB := config.GetDB() // 初始化数据库句柄

	getTeamC.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		team := ParseMatchTeam(dom)
		team.TeamUrl = teamUrl
		OperateMatchTeam(DB, team)

		if len(team.Players) > 0 {
			for _, player := range team.Players {
				scrapyPlayer(player.PlayerUrl)
			}
		}
	})

	getTeamC.OnResponse(func(r *colly.Response) {
		fmt.Println("访问战队网页 Visited ", r.Request.URL.String())
	})

}

func ParseMatchTeam(dom *goquery.Document) model.Team {
	//解析比赛战队网页
	var team model.Team
	playerResultSet := make([]model.Player, 0)

	teamProfileDom := dom.Find("div[class='profile-team-stats-container']").Find("div[class='profile-team-stat']")
	worldRankingStr := teamProfileDom.Eq(0).Find("span").Text()
	worldRanking, _ := strconv.Atoi(strings.Replace(worldRankingStr, "#", "", -1))
	averagePlayerAgeStr := teamProfileDom.Eq(2).Find("span").Text()
	averagePlayerAge, _ := strconv.ParseFloat(averagePlayerAgeStr, 64)
	averagePlayerAge = utils.Decimal(averagePlayerAge)
	teamName := dom.Find("div[class='profile-team-info']").Find("h1[class='profile-team-name text-ellipsis']").Text()
	coatchName := teamProfileDom.Eq(3).Find("span").Text()
	coatchName = strings.Replace(coatchName, "'", "", -1)

	profileTopDom := dom.Find("div[class='standard-box profileTopBox clearfix']").Find("div[class='flex']")
	teamPic, _ := profileTopDom.Find("div[class='profile-team-logo-container']").Find("img").Attr("src")
	nationName := profileTopDom.Find("div[class='profile-team-info']").Find("div[class='team-country text-ellipsis']").Text()
	nationPic, _ := profileTopDom.Find("div[class='profile-team-info']").Find("div[class='team-country text-ellipsis']").Find("img").Attr("src")
	nationPic = parameter.HLTV_INDEX + nationPic

	team.TeamName = utils.CompressString(teamName)
	team.WorldRanking = uint16(worldRanking)
	team.AveragePlayerAge = float32(averagePlayerAge)
	team.CoatchName = utils.CompressString(coatchName)
	team.TeamPic = teamPic
	team.NationName = utils.CompressString(nationName)
	team.NationPic = nationPic

	//fmt.Println("teamPic=", teamPic)
	//fmt.Println("nationName=", nationName)
	//fmt.Println("nationPic=", nationPic)
	//fmt.Println("teamName=", teamName)
	//fmt.Printf("worldRanking=%v,%T\n", worldRanking, worldRanking)
	//fmt.Printf("averagePlayerAge=%v,%T\n", averagePlayerAge, averagePlayerAge)
	//fmt.Println("coatchName=", coatchName)

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
