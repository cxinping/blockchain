package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strconv"
	"strings"
)

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
