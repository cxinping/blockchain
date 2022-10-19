package main

import (
	"spider/src/config"
)

func main() {
	//config.InitConfig() //读取配置文件
	//fmt.Println("读取配置文件config.yml 得到 db.host => ", viper.Get("db.host"))

	db := config.InitDB() // 初始化数据库
	defer db.Close()

	//crawl.CrawlTournamentWeb() //爬取赛事和比赛数据
	//task.ExecTasks() //执行调度任务

	//start := time.Now()
	//crawl.CrawlMatcheResultWeb("https://www.hltv.org/results")
	//elapsed := time.Since(start)
	//fmt.Printf("WaitGroupStart Time %s\n ", elapsed)
}
