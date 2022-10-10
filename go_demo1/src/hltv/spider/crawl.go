package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CrawlMatches() {
	// 爬取赛事信息
	base_url := "https://www.hltv.org/matches"
	fmt.Println("*** 开始爬取hltv的赛事列表 ", base_url)

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	// selector goquery name id class
	c.OnHTML(".upcoming-headline", func(e *colly.HTMLElement) {
		fmt.Println("读取class为upcoming-headline的数据", e.Text)
		//bodyData := e.Response.Body
		//fmt.Println(string(bodyData))

		//e.Request.Visit(e.Attr("href"))
		//
		//ret, _ := e.DOM.Html()
		//fmt.Println("ret-> ", ret)
	})

	c.OnHTML(".upcoming-headline", func(e *colly.HTMLElement) {
		fmt.Println("读取class为upcoming-headline的数据", e.Text)

	})

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", RandomString())

		fmt.Println("OnRequest")
		//fmt.Println("url => ", r.URL)
	})

	c.Visit(base_url)

}
