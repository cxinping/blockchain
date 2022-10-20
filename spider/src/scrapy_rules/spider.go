package scrapy_rules

import "fmt"

func ScrapyPlayerInfomation(playerUrl string) {
	//爬取队员信息
	getPlayerC := GetDefaultCollector()
	SetPlayerCallback(getPlayerC, playerUrl)

	err := getPlayerC.Visit(playerUrl)
	if err != nil {
		fmt.Println("访问网页", playerUrl, "具体错误:", err)
	}

	getPlayerC.Wait()
}
