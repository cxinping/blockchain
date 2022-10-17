package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strings"
)

func ParseMatchDetail(dom *goquery.Document) (string, []model.Team, []model.Player) {
	//解析比赛网页数据, 抓取战队数据
	fmt.Println("*** OperateMatchDetail ***", dom)
	teamResultSet := make([]model.Team, 0)
	playResultSet := make([]model.Player, 0)
	var team1 model.Team
	var team2 model.Team
	teamBoxDom := dom.Find("div[class='standard-box teamsBox']")
	team1Name := teamBoxDom.Find("div[class='teamName']").Eq(0).Text()
	team2Name := teamBoxDom.Find("div[class='teamName']").Eq(1).Text()
	team1Dom := teamBoxDom.Find("div[class='team1-gradient']")
	team2Dom := teamBoxDom.Find("div[class='team2-gradient']")
	team1Pic, _ := team1Dom.Find("img").Attr("src")
	team2Pic, _ := team2Dom.Find("img").Attr("src")
	team1Url, _ := team1Dom.Find("a").Attr("href")
	team2Url, _ := team2Dom.Find("a").Attr("href")

	fmt.Println("team1Name=", team1Name, ", team2Name=", team2Name)
	//fmt.Println("team1Pic=", team1Pic)
	//fmt.Println("team2Pic=", team2Pic)
	fmt.Println("team1Url=", parameter.HLTV_INDEX+team1Url)
	fmt.Println("team2Url=", parameter.HLTV_INDEX+team2Url)

	team1.TeamName = team1Name
	team1.TeamPic = team1Pic
	team2.TeamName = team2Name
	team2.TeamPic = team2Pic

	mapsStr := dom.Find("div[class='padding preformatted-text']").Text()
	mapsStr = utils.CompressString(strings.ToLower(mapsStr))
	count := strings.Index(mapsStr, "online")
	var modeStr string = "" // 比赛是线上还是线下
	if count != -1 {
		modeStr = parameter.MATCH_MODE_ONLINE
	} else {
		modeStr = parameter.MATCH_MODE_LAN
	}

	//fmt.Println("modeStr=>", modeStr)

	return modeStr, teamResultSet, playResultSet
}
