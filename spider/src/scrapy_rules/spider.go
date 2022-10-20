package scrapy_rules

import (
	"fmt"
	"spider/src/utils/parameter"
)

func ScrapyPlayerInfomation(playerUrl string) {
	//爬取队员数据
	getPlayerC := GetDefaultCollector()
	SetPlayerCallback(getPlayerC, playerUrl)

	err := getPlayerC.Visit(playerUrl)
	if err != nil {
		fmt.Println("访问网页", playerUrl, "具体错误:", err)
	}

	getPlayerC.Wait()
}

func ScrapyTeamInformation(teamUrl string) {
	//爬取战队数据
	getTeamC := GetDefaultCollector()
	SetTeamCallback(getTeamC, teamUrl)

	err := getTeamC.Visit(teamUrl)
	if err != nil {
		fmt.Println("访问网页", teamUrl, "具体错误:", err)
	}

	getTeamC.Wait()
}

func ScrapyMatchInformation(matchUrl string) {
	//爬取比赛数据
	getMatchC := GetDefaultCollector()
	SetMatchCallback(getMatchC, matchUrl, ScrapyTeamInformation)
	err := getMatchC.Visit(matchUrl)
	if err != nil {
		fmt.Println("访问网页", matchUrl, "具体错误:", err)
	}

	getMatchC.Wait()
}

func ScrapyMatchResults() {
	// 爬取已经有结果的赛果数据
	getMatchC := GetDefaultCollector()
	SetMatcheResults(getMatchC)

	err := getMatchC.Visit(parameter.MATCH_RESULT_URL)
	if err != nil {
		fmt.Println("访问网页", parameter.MATCH_RESULT_URL, "具体错误:", err)
	}

	getMatchC.Wait()
}
