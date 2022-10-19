package main

import (
	"fmt"
	"github.com/spf13/viper"
	"spider/src/config"
)

func main() {
	config.InitConfig() //读取配置文件
	fmt.Println("读取配置文件config.yml 得到 db.host => ", viper.Get("db.host"))

	db := config.InitDB() // 初始化数据库
	defer db.Close()

	//crawl.CrawlTournamentWeb() //爬取赛事和比赛数据
	//task.ExecTasks() //执行调度任务

}
