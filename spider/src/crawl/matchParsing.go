package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strconv"
	"strings"
	"time"
)

func ParseMatchDetail(dom *goquery.Document) (time.Time, string, string, model.Team, model.Team) {
	//解析比赛网页数据, 抓取战队数据
	//fmt.Println("*** OperateMatchDetail ***")

	var team1 model.Team
	var team2 model.Team
	play1ResultSet := make([]model.Player, 0)
	play2ResultSet := make([]model.Player, 0)

	teamBoxDom := dom.Find("div[class='standard-box teamsBox']")
	team1Name := teamBoxDom.Find("div[class='teamName']").Eq(0).Text()
	team2Name := teamBoxDom.Find("div[class='teamName']").Eq(1).Text()
	team1Dom := teamBoxDom.Find("div[class='team1-gradient']")
	team2Dom := teamBoxDom.Find("div[class='team2-gradient']")
	team1Pic, _ := team1Dom.Find("img").Attr("src")
	team2Pic, _ := team2Dom.Find("img").Attr("src")
	team1Url, _ := team1Dom.Find("a").Attr("href")
	team2Url, _ := team2Dom.Find("a").Attr("href")

	//fmt.Println("team1Name=", team1Name, ", team2Name=", team2Name)
	//fmt.Println("team1Pic=", team1Pic)
	//fmt.Println("team2Pic=", team2Pic)
	//fmt.Println("team1Url=", parameter.HLTV_INDEX+team1Url)
	//fmt.Println("team2Url=", parameter.HLTV_INDEX+team2Url)

	team1.TeamName = team1Name
	team1.TeamPic = team1Pic
	team1.TeamUrl = parameter.HLTV_INDEX + team1Url

	team2.TeamName = team2Name
	team2.TeamPic = team2Pic
	team2.TeamUrl = parameter.HLTV_INDEX + team2Url

	mapsStr := dom.Find("div[class='padding preformatted-text']").Text()
	mapsStr = utils.CompressString(strings.ToLower(mapsStr))
	count := strings.Index(mapsStr, "online")
	var matchMode string = "" // 比赛模式是线上还是线下
	if count != -1 {
		matchMode = parameter.MATCH_MODE_ONLINE
	} else {
		matchMode = parameter.MATCH_MODE_LAN
	}
	matchDateUnixStr, _ := dom.Find("div[class='timeAndEvent']").Find("div[class='date']").Attr("data-unix")
	matchDateUnixInt, _ := strconv.ParseInt(matchDateUnixStr, 10, 64)
	matchDateUnixInt = int64(matchDateUnixInt) / 1000
	matchTime := time.Unix(matchDateUnixInt, 0)                                                   // 比赛时间
	matchStatusStr := dom.Find("div[class='timeAndEvent']").Find("div[class='countdown']").Text() //比赛状态
	matchStatus := parameter.MATCH_STATUS_UNSTARTED
	//fmt.Println("* matchDateUnixStr=", matchDateUnixStr)
	//fmt.Println("* matchTime=", matchTime)

	//fmt.Println("* matchMode=", matchMode)
	if matchStatusStr == "Match over" {
		matchStatus = parameter.MATCH_STATUS_OVER
	} else if matchStatusStr == "LIVE" {
		matchStatus = parameter.MATCH_STATUS_LIVE
	}
	//fmt.Println("* matchStatusStr=", matchStatusStr, ",matchStatus=", matchStatus)

	// 战队1的所有队员
	player1Dom := dom.Find("div[class='lineup standard-box']").Eq(0)
	player1Dom.Find("td[class='player player-image']").Each(func(idx int, tdSel *goquery.Selection) {
		player := model.Player{}
		playerPic, _ := tdSel.Find("img").Attr("src")
		playerTdDom := player1Dom.Find("td[class='player']").Eq(idx)
		playerName := playerTdDom.Find("div[class='text-ellipsis']").Text()
		nationPic, _ := playerTdDom.Find("img").Attr("src")
		nationPic = parameter.HLTV_INDEX + nationPic
		player.PlayerPic = playerPic
		player.NationPic = nationPic
		player.Name = playerName
		//fmt.Println("idx=", idx, ", playerPic=", playerPic, ",playerName=", playerName, ",nationPic=", nationPic)
		play1ResultSet = append(play1ResultSet, player)
	})
	team1.Players = play1ResultSet

	// 战队2的所有队员
	player2Dom := dom.Find("div[class='lineup standard-box']").Eq(1)
	player2Dom.Find("td[class='player player-image']").Each(func(idx int, tdSel *goquery.Selection) {
		player := model.Player{}
		playerPic, _ := tdSel.Find("img").Attr("src")
		playerTdDom := player2Dom.Find("td[class='player']").Eq(idx)
		playerName := playerTdDom.Find("div[class='text-ellipsis']").Text()
		nationPic, _ := playerTdDom.Find("img").Attr("src")
		nationPic = parameter.HLTV_INDEX + nationPic
		player.PlayerPic = playerPic
		player.NationPic = nationPic
		player.Name = playerName
		//fmt.Println("idx=", idx, ", playerPic=", playerPic, ",playerName=", playerName, ",nationPic=", nationPic)
		play2ResultSet = append(play2ResultSet, player)
	})
	team2.Players = play2ResultSet

	return matchTime, matchMode, matchStatus, team1, team2
}
