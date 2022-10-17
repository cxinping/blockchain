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
		link := e.Attr("href")
		link = "https://gorm.io/zh_CN/docs/" + link
		fmt.Printf("Link found: %q -> %s\n", e.Text, "https://gorm.io/zh_CN/docs/"+link)

		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("访问网页 ", r.URL.String())
	})

	c.Visit("https://gorm.io/zh_CN/docs")
}
