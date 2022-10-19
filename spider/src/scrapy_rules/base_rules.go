package scrapy_rules

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"os"
	"time"
)

// return a collector
func GetDefaultCollector() *colly.Collector {
	//set async and dont forget set c.wait()
	debugger := &debug.LogDebugger{}

	file, err := os.Create("/Users/xinping/topgaming/debug.log")
	if err != nil {
		panic(err)
	}
	debugger.Output = file

	var c = colly.NewCollector(
		colly.Async(true),
		colly.Debugger(debugger),
	)
	//disable http KeepAlives its could cause OOM in long time work
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	setDefaultCallback(c)
	extensions.RandomUserAgent(c)
	return c
}

// set default call,cookie and error handling
func setDefaultCallback(c *colly.Collector) {
	// set random cookie
	c.OnRequest(func(r *colly.Request) {

	})

	// Limit the maximum parallelism to 2
	// This is necessary if the goroutines are dynamically
	// created to control the limit of simultaneous requests.
	//
	// Parallelism can be controlled also by spawning fixed
	// number of go routines.

	// delay 3 to 5 second
	delay := time.Duration(5)
	randomDelay := time.Duration(5)
	if delay == 0 || randomDelay == 0 {
		delay, randomDelay = 8, 2
	}
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2, Delay: delay * time.Second, RandomDelay: randomDelay * time.Second})

	// deal with error statusCode
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err, "\nStatusCode", r.StatusCode)

	})
}
