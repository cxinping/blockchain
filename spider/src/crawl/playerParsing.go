package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strconv"
	"strings"
)

func ParseMatchTeamPlayer(dom *goquery.Document) model.Player {
	//解析比赛战队的队员网页
	var player model.Player
	playerContainerDom := dom.Find("div[class='playerContainer']")
	playerPic, _ := playerContainerDom.Find("img").Eq(1).Attr("src")
	nickName := playerContainerDom.Find("h1[class='playerNickname']").Text()
	nationPic, _ := playerContainerDom.Find("div[class='playerRealname']").Find("img").Attr("src")
	nationPic = parameter.HLTV_INDEX + nationPic
	nationName := playerContainerDom.Find("div[class='playerRealname']").Text()
	ageStr := playerContainerDom.Find("div[class='playerAge']").Find("span[class='listRight']").Find("span").Text()
	ageStr = utils.CompressString(strings.Replace(ageStr, "years", "", -1))
	age, _ := strconv.Atoi(ageStr)
	currentTeamPic, _ := playerContainerDom.Find("div[class='playerTeam']").Find("span[class='listRight']").Find("img").Attr("src")
	currentTeamName := playerContainerDom.Find("div[class='playerTeam']").Find("span[class='listRight']").Find("a").Text()

	//fmt.Println("playerPic=", playerPic)
	//fmt.Println("nickName=", nickName)
	//fmt.Println("nationPic=", nationPic)
	//fmt.Println("nationName=", nationName)
	//fmt.Printf("age=[%v],%T\n", age, age)
	fmt.Printf("currentTeamPic=[%v],%T\n", currentTeamPic, currentTeamPic)
	fmt.Printf("currentTeamName=[%v],%T\n", currentTeamName, currentTeamName)

	player.PlayerPic = playerPic
	player.NickName = utils.CompressString(nickName)
	player.NationPic = nationPic
	player.NationName = utils.CompressString(nationName)
	player.PlayerAge = uint8(age)
	player.CurrentTeamName = utils.CompressString(currentTeamName)
	player.CurrentTeamPic = currentTeamPic

	return player
}
