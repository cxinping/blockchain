package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"math/rand"
	"strconv"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CompressString(input_str string) string {
	input_str = strings.Replace(input_str, "\n", "", -1)
	input_str = strings.Trim(input_str, " ")
	return input_str
}

func CrawlMatches() {
	// 爬取赛事信息
	base_url := "https://www.hltv.org/matches"
	fmt.Println("*** 开始爬取hltv的赛事列表 ", base_url)

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	//c.OnHTML(".upcoming-headline", func(e *colly.HTMLElement) {
	//	fmt.Println("读取class为upcoming-headline的数据", e.Text)
	//	bodyData := e.Response.Body
	//	fmt.Println(string(bodyData))
	//
	//	//e.Request.Visit(e.Attr("href"))
	//	//ret, _ := e.DOM.Html()
	//	//fmt.Println("ret-> ", ret)
	//})

	//c.OnHTML(".upcomingMatchesSection", func(e *colly.HTMLElement) {
	//	fmt.Println("读取class为upcomingMatchesSection的数据")
	//
	//})

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", RandomString())
		//fmt.Println("OnRequest")
		fmt.Println("url => ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("OnResponse")
		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))

		operate_living_match(dom)

		operate_upcoming_match(dom)

	})

	c.Visit(base_url)

}

func operate_upcoming_match(dom *goquery.Document) {
	/**
	处理预计比赛的数据
	*/
	//upcoming_match := live_match_section_dom.Next().Text()
	upcoming_match := dom.Find(".upcoming-headline").Text()
	upcoming_match = strings.Replace(upcoming_match, "\n", "", -1)
	upcoming_match = strings.Trim(upcoming_match, " ")
	fmt.Println("预计比赛的赛事名称=", upcoming_match)

	// .upcomingMatchesAll .upcomingMatchesSection
	dom.Find(".upcomingMatchesSection").Each(func(i int, selection *goquery.Selection) {
		// 比赛时间
		match_date := selection.Find(".matchDayHeadline").Text()
		//match_date = strings.Replace(match_date, " ", "", -1)
		//match_date_idx := strings.Index(match_date, "-") + 1
		//match_date = string([]rune(match_date)[match_date_idx:len(match_date)])
		fmt.Println("idx=>", i+1, ",match_date=", match_date)

		selection.Find(".upcomingMatch").Each(func(i int, selection *goquery.Selection) {
			match_time := selection.Find(".matchInfo .matchTime").Text()
			fmt.Println("\tmatch_time=", match_time)

			match_date_unix_str, _ := selection.Find(".matchInfo .matchTime").Attr("data-unix")
			match_date_unix_int, _ := strconv.ParseInt(match_date_unix_str, 10, 64)
			match_date_unix_int = int64(match_date_unix_int) / 1000
			//match_time := time.Unix(match_date_unix_int, 0).Format("2006-01-02 15:04")
			//fmt.Println("\tmatch_time=", match_time)
			team1_name := CompressString(selection.Find("div[class='matchTeam team1']").Text())
			team1_pic, _ := selection.Find("div[class='matchTeam team1']").Find("img").Attr("src")
			team2_name := CompressString(selection.Find("div[class='matchTeam team2']").Text())
			team2_pic, _ := selection.Find("div[class='matchTeam team2']").Find("img").Attr("src")
			match_pic, _ := selection.Find(".matchEvent").Find(".matchEventLogoContainer").Find("img").Attr("src") // 比赛的图片logo
			match_name := selection.Find("div[class='matchEventName gtSmartphone-only']").Text()
			fmt.Println("\tteam1_name=", team1_name)
			fmt.Println("\tteam1_pic=", team1_pic)
			fmt.Println("\tteam2_name=", team2_name)
			fmt.Println("\tteam2_pic=", team2_pic)
			fmt.Println("\tmatch_pic=", match_pic)
			fmt.Println("\tmatch_name=", match_name)
			fmt.Println("")
		})
		fmt.Println("")
	})
}

func operate_living_match(dom *goquery.Document) {
	/**
	处理正赛比赛的数据
	*/
	live_match_section_dom := dom.Find(".liveMatchesSection")
	if live_match_section_dom != nil {
		// 赛事名称
		match_name := live_match_section_dom.Find(".upcoming-headline").Text()
		fmt.Printf("正在比赛的赛事名称=%v\n", match_name)

		live_match_section_dom.Find(".liveMatch-container").Each(func(i int, selection *goquery.Selection) {
			match_link, _ := selection.Find("a[class='match a-reset']").Attr("href")
			fmt.Println("正在比赛的地址=> ", "https://www.hltv.org"+match_link)

			selection.Find("div[class='matchTeam']").Each(func(i int, selection *goquery.Selection) {
				team_name := CompressString(selection.Find("div[class='matchTeamName text-ellipsis']").Text())
				if i == 0 {
					team1_name := team_name
					fmt.Println("\tteam1_name=", team1_name)
				} else if i == 1 {
					team2_name := team_name
					fmt.Println("\tteam2_name=", team2_name)
				}

			})
			fmt.Println("")
		})

		fmt.Println()
	}
}
