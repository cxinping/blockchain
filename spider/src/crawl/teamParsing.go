package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils/parameter"
	"strconv"
	"strings"
)

func ParseMatchTeam(dom *goquery.Document) model.Team {
	//解析比赛战队网页
	var team model.Team
	playerResultSet := make([]model.Player, 0)

	dom.Find("div[class='bodyshot-team g-grid']").Find("a[class='col-custom']").Each(func(idx int, selection *goquery.Selection) {
		var player model.Player
		playerUrl, _ := selection.Attr("href")
		playerUrl = parameter.HLTV_INDEX + playerUrl
		player.PlayerUrl = playerUrl
		playerPic, _ := selection.Find("img").Attr("src")
		playerName := selection.Find("div[class='playerFlagName']").Find("span[class='text-ellipsis bold']").Text()
		player.PlayerPic = playerPic
		player.Name = playerName
		nationPic, _ := selection.Find("span[class='gtSmartphone-only']").Find("img").Attr("src")
		nationPic = parameter.HLTV_INDEX + nationPic
		player.NationPic = nationPic

		//fmt.Println(idx, playerName, nationPic)
		playerResultSet = append(playerResultSet, player)
	})
	team.Players = playerResultSet
	teamProfileDom := dom.Find("div[class='profile-team-stats-container']").Find("div[class='profile-team-stat']")
	worldRankingStr := teamProfileDom.Eq(0).Find("span").Text()
	worldRanking, _ := strconv.Atoi(strings.Replace(worldRankingStr, "#", "", -1))
	averagePlayerAge := teamProfileDom.Eq(2).Find("span").Text()
	coatchName := teamProfileDom.Eq(3).Find("span").Text()
	coatchName = strings.Replace(coatchName, "'", "", -1)

	//team.WorldRanking = worldRanking

	fmt.Printf("%v,%T\n", worldRanking, worldRanking)
	fmt.Println("averagePlayerAge=", averagePlayerAge)
	fmt.Println("coatchName=", coatchName)

	return team
}
