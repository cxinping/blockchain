package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils"
)

func ParseMatchTeamPlayer(dom *goquery.Document) model.Player {
	//解析比赛战队的队员网页
	var player model.Player
	playerContainerDom := dom.Find("div[class='playerContainer']")
	playerPic, _ := playerContainerDom.Find("img").Eq(1).Attr("src")
	nickName := playerContainerDom.Find("h1[class='playerNickname']").Text()

	fmt.Println("playerPic=", playerPic)
	fmt.Println("nickName=", nickName)

	player.PlayerPic = playerPic
	player.NickName = utils.CompressString(nickName)

	return player
}
