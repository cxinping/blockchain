package main

import (
	"spider/src/config"
	"spider/src/crawl"
	"spider/src/task"
)

func main() {
	db := config.InitDB() // 初始化数据库
	defer db.Close()

	crawl.CrawlMatches() //爬取赛事和比赛数据
	task.QueryMatches()  //执行调度任务
}
