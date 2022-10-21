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
	// 爬取多页的赛果数据
	getMatchC := GetDefaultCollector()
	SetMatcheResults(getMatchC, ScrapyMatchResult)

	err := getMatchC.Visit(parameter.MATCH_RESULT_URL)
	if err != nil {
		fmt.Println("访问网页", parameter.MATCH_RESULT_URL, "具体错误:", err)
	}

	getMatchC.Wait()
}

func ScrapyMatchResult(matchURL string) {
	// 爬取每一页比赛结果列表的比赛数据
	getMatchC := GetDefaultCollector()
	SetMatcheResult(getMatchC, ScrapyMatchInformation)

	err := getMatchC.Visit(matchURL)
	if err != nil {
		fmt.Println("访问网页", matchURL, "具体错误:", err)
	}

	getMatchC.Wait()
}

/////////////////

func ScrapyTournament() {
	getTournamentC := GetDefaultCollector()

	SetTournamentCallback(getTournamentC)
	err := getTournamentC.Visit(parameter.TT_MATCH_URL)
	if err != nil {
		fmt.Println("访问网页", parameter.TT_MATCH_URL, "具体错误:", err)
	}

	getTournamentC.Wait()
}
