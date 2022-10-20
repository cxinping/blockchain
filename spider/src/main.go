package main

import (
	"fmt"
	"github.com/spf13/viper"
	"spider/src/config"
	"spider/src/scrapy_rules"
)

func init() {
	config.InitConfig() //读取配置文件
	fmt.Println("读取配置文件config.yml 得到 db.host => ", viper.Get("db.host"))

	db := config.InitDB() // 初始化数据库
	defer db.Close()
}

func scrapyPlayerWeb() {
	getPlayerC := scrapy_rules.GetDefaultCollector()
	//fmt.Println("getPlayerC=", getPlayerC)
	scrapy_rules.SetPlayerCallback(getPlayerC, "")

	getPlayerC.Wait()
}

func main() {

	scrapyPlayerWeb()
	//crawl.CrawlTournamentWeb() //爬取赛事和比赛数据
	//task.ExecTasks() //执行调度任务

	//start := time.Now()
	//crawl.CrawlMatcheResultWeb("https://www.hltv.org/results")
	//elapsed := time.Since(start)
	//fmt.Printf("WaitGroupStart Time %s\n ", elapsed)
}
