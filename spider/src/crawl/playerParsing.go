package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
)

func ParseMatchTeamPlayer(dom *goquery.Document) model.Player {
	//解析比赛战队的队员网页
	var player model.Player
	playerContainerDom := dom.Find("div[class='playerContainer']")
	playerPic, _ := playerContainerDom.Find("img").Eq(1).Attr("src")

	fmt.Println(playerPic)

	player.PlayerPic = playerPic
	return player
}
