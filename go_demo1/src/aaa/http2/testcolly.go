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

		//ret, _ := e.DOM.Html()
		//fmt.Println("ret-> ", ret)
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
		fmt.Printf("Link found: %q -> %s\n", e.Text, "https://gorm.io/zh_CN/docs/"+link)
		c.Visit(e.Request.AbsoluteURL(link))

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://gorm.io/zh_CN/docs")
}
