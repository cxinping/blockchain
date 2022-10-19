package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"net"
	"net/http"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strings"
	"time"
)

// 爬取将要进行的比赛网页数据, 从 https://www.hltv.org/matches 获得的比赛页面 url
func CrawlMatcheWeb(matchUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

	// 爬取赛事信息
	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		//colly.Async(true),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second, // 超时时间
			KeepAlive: 300 * time.Second, // keepAlive 超时时间
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 10 * time.Second,
	})
	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	//爬取赛事和赛果网页数据
	c.OnHTML("div.match-page", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

		match, team1, team2 := ParseMatchDetail(dom)
		match.MatchUrl = matchUrl
		operateMatchDetail(DB, match, team1, team2)
		CrawlTeam(team1.TeamUrl)
		CrawlTeam(team2.TeamUrl)

		//matchTime, matchMode, matchStatus, team1, team2 := ParseMatchDetail(dom)
		////fmt.Println("matchTime=", matchTime)
		////fmt.Println("matchModeStr=", matchModeStr)
		////fmt.Println("team1.TeamUrl=", team1.TeamUrl)
		////fmt.Println(team2)
		//operateMatchDetail(DB, matchUrl, matchTime, matchMode, matchStatus, team1, team2)
		//CrawlTeam(team1.TeamUrl)
		//CrawlTeam(team2.TeamUrl)
	})

	// 异常处理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问比赛网页 Visited ", r.Request.URL.String())
	})

	c.Visit(matchUrl)
}

func CrawlMatcheResults() {
	// 爬取已经有结果的赛果数据
	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		//colly.Async(true),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second, // 超时时间
			KeepAlive: 300 * time.Second, // keepAlive 超时时间
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 10 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	// 异常处理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("div[class='results-holder allres']", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		requestUrls := ParseMatchPageOffset(dom)

		for _, reqeustUrl := range requestUrls {
			fmt.Println("111 reqeustUrl=", reqeustUrl)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问已经有比赛结果的赛果网页 Visited ", r.Request.URL.String())
	})

	c.Visit(parameter.MATCH_RESULT_URL)
}

// 爬取已经结束的比赛网页数据
func CrawlMatcheResultWeb(matchUrl string) {
	//分页爬取比赛结果的网页数据

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		//colly.Async(true),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   300 * time.Second, // 超时时间
			KeepAlive: 300 * time.Second, // keepAlive 超时时间
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 10 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	c.OnHTML("div[class='results-all']", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))

		matchUrls := ParseMatchResult(dom)
		// 方法1
		for _, matchUrl := range matchUrls {
			//fmt.Println("matchUrl=> ", matchUrl)
			CrawlMatcheWeb(matchUrl)

		}

		// 方法2
		//wg := sync.WaitGroup{}
		//
		//for _, matchUrl := range matchUrls {
		//	wg.Add(1)
		//
		//	go func() {
		//		CrawlMatcheWeb(matchUrl)
		//		time.Sleep(2 * time.Second)
		//
		//		wg.Done()
		//	}()
		//}
		//wg.Wait()

		fmt.Printf("一共处理 %d 条页面的比赛页面", len(matchUrls))
	})

	// 异常处理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问已经有比赛结果的赛果网页 Visited ", r.Request.URL.String())
	})

	c.Visit(matchUrl)
}

func operateMatchDetail(DB *gorm.DB, match model.Match, team1 model.Team, team2 model.Team) {
	//处理比赛详细数据
	//var queryMatch = model.Match{}
	//DB.Where("match_url = ?", match.MatchUrl).Find(&queryMatch)
	//// 修改比赛的状态
	//if match.Status == parameter.MATCH_STATUS_LIVE {
	//	DB.Model(&match).Updates(model.Match{Mode: match.Mode, MatchTime: match.MatchTime, Status: match.Status})
	//} else if match.Status == parameter.MATCH_STATUS_UNSTARTED {
	//	DB.Model(&match).Updates(model.Match{Mode: match.Mode, Status: match.Status})
	//}

	// 处理战队数据
	var count int = 0
	DB.Model(&model.Team{}).Where("team_name = ?", team1.TeamName).Count(&count)
	if count == 0 {
		team1.TeamBizId = utils.GenerateModuleBizID("TM")
		team1.CreatedTime = time.Now()
		team1.Insert(DB)
	}
	count = 0
	DB.Model(&model.Team{}).Where("team_name = ?", team2.TeamName).Count(&count)
	if count == 0 {
		team2.TeamBizId = utils.GenerateModuleBizID("TM")
		team2.CreatedTime = time.Now()
		team2.Insert(DB)
	}

	// 处理比赛数据
	matchCount := 0
	DB.Model(&model.Match{}).Where("match_url = ?", match.MatchUrl).Count(&matchCount)
	if matchCount == 0 {
		//插入match对象，包括相关的2个team对象的属性
		//fmt.Println("111 insert match record")
		tt := model.Tournament{}
		DB.Where("tt_name = ?", match.TtName).Find(&tt)

		queryTeam1 := model.Team{}
		DB.Where("team_url = ?", team1.TeamUrl).Find(&queryTeam1)
		queryTeam2 := model.Team{}
		DB.Where("team_url = ?", team2.TeamUrl).Find(&queryTeam2)

		match.MatchBizId = utils.GenerateModuleBizID("MH")
		match.CreatedTime = time.Now()
		match.TtBizId = tt.TtBizId
		match.TtPic = tt.TtPic
		match.Team1Name = team1.TeamName
		if queryTeam1.TeamBizId != "" {
			match.Team1BizId = queryTeam1.TeamBizId
		} else {
			match.Team1BizId = team1.TeamBizId
		}

		match.Team2Name = team2.TeamName
		if queryTeam2.TeamBizId != "" {
			match.Team2BizId = queryTeam2.TeamBizId
		} else {
			match.Team2BizId = team2.TeamBizId
		}

		match.Insert(DB)
	} else {
		//fmt.Println("222 update match record")
		DB.Model(model.Match{}).Where("match_url = ?", match.MatchUrl).Update(match)
	}

}
