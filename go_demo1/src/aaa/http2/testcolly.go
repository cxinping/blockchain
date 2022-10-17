package http2

import (
	"fmt"
	"github.com/gocolly/colly"
)

func Example1() {
	c := colly.NewCollector()

	// selector goquery name id class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		content, _ := e.DOM.Html()
		link := e.Attr("href")
		link = "https://gorm.io/zh_CN/docs/" + link
		fmt.Println("content=", content, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url => ", r.URL)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://gorm.io/zh_CN/docs")
}

func Example2() {
	c := colly.NewCollector()

	// selector goquery name id class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		fmt.Println("OnResponse收到html内容后调用:OnHTML ")

		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		fmt.Println("")
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("获得响应后调用:OnResponse")
		//fmt.Println("Visited", r.Request.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("访问网页 ", r.URL.String())
	})

	c.Visit("https://gorm.io/zh_CN/docs")
}

func Example3() {
	c := colly.NewCollector()

	// selector goquery name id class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))

		//bodyData := string(r.Body)
		//dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))

		fmt.Println("OnResponse收到html内容后调用:OnHTML ")

		link := e.Attr("href")
		link = e.Request.AbsoluteURL(link)

		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		fmt.Println("")
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("获得响应后调用:OnResponse")
		//fmt.Println("Visited", r.Request.URL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("访问网页 ", r.URL.String())
	})

	c.Visit("https://gorm.io/zh_CN/docs")
}
