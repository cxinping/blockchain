package main

import (
	"spider/src/config"
	"spider/src/crawl"
)

func main() {
	db := config.InitDB() // 初始化数据库
	defer db.Close()

	crawl.CrawlMatches()
}
