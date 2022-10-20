package scrapy_rules

func ScrapyPlayerInfomation(playerUrl string) {
	//爬取队员信息
	getPlayerC := GetDefaultCollector()
	SetPlayerCallback(getPlayerC, playerUrl)

	getPlayerC.Visit(playerUrl)
	getPlayerC.Wait()
}
