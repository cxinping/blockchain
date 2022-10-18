package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils/parameter"
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

	return team
}
