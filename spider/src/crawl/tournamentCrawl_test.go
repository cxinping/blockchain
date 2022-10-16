package crawl

import (
	"spider/src/config"
	"testing"
)

func init() {
	config.InitDB()
}

func TestCrawlMatches(t *testing.T) {
	t.Log("*** 开始解析赛事网页 ***")
	CrawlMatches()
}
