package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"strings"
	"time"
)

func CrawlTeamHttp(teamUrl string) {
	resp, err := http.Get(teamUrl)
	if err != nil {
		fmt.Print("http get err", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("http get err", err)
		return
	}

	fmt.Println(string(body))
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	team := ParseMatchTeam(dom)

	team.TeamUrl = teamUrl
	DB := config.GetDB() // 初始化数据库句柄
	OperateMatchTeam(DB, team)

	// 方法1
	if len(team.Players) > 0 {
		for _, player := range team.Players {
			//fmt.Println("player.PlayerUrl=> ", player.PlayerUrl)
			CrawlPlayerHttp(player.PlayerUrl)
		}
	}

	// 方法2
	//wg := sync.WaitGroup{}
	//
	//for _, player := range team.Players {
	//	wg.Add(1)
	//
	//	go func() {
	//		CrawlPlayerHttp(player.PlayerUrl)
	//		time.Sleep(3 * time.Second)
	//
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()

}

// 爬取战队的网页数据
func CrawlTeam(teamUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

	c := colly.NewCollector(
		colly.AllowedDomains("www.hltv.org"), //白名单域名
		// 允许重复访问
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		colly.Async(true),
		//colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.MaxDepth(1),            //爬取页面深度,最多为两层
		colly.MaxBodySize(1024*1024), //响应正文最大字节数
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"),
		colly.IgnoreRobotsTxt(), //忽略目标机器中的`robots.txt`声明
	)
	c.SetRequestTimeout(120 * time.Second)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       1 * time.Second,
		RandomDelay: 5 * time.Second,
	})

	//c.WithTransport(&http.Transport{
	//	Proxy: http.ProxyFromEnvironment,
	//	DialContext: (&net.Dialer{
	//		Timeout:   300 * time.Second, // 超时时间
	//		KeepAlive: 300 * time.Second, // keepAlive 超时时间
	//		DualStack: true,
	//	}).DialContext,
	//	MaxIdleConns:          100,              // 最大空闲连接数
	//	IdleConnTimeout:       90 * time.Second, // 空闲连接超时
	//	TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
	//	ExpectContinueTimeout: 10 * time.Second,
	//})

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", utils.RandomString())
	})

	c.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		team := ParseMatchTeam(dom)
		team.TeamUrl = teamUrl
		OperateMatchTeam(DB, team)

		// 方法1
		if len(team.Players) > 0 {
			for _, player := range team.Players {
				//fmt.Println("player.PlayerUrl=> ", player.PlayerUrl)
				CrawlPlayer(player.PlayerUrl)
			}
		}

		//方法2
		//wg := sync.WaitGroup{}
		//for _, player := range team.Players {
		//	wg.Add(1)
		//
		//	go func() {
		//		time.Sleep(time.Millisecond * 20)
		//		CrawlPlayerHttp(player.PlayerUrl)
		//		//time.Sleep(2 * time.Second)
		//
		//		wg.Done()
		//	}()
		//}
		//wg.Wait()

	})

	// 异常处理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问战队网页 Visited ", r.Request.URL.String())
	})

	err := c.Visit(teamUrl)
	if err != nil {
		fmt.Println("具体错误:", err)
	}

	c.Wait()
}

func OperateMatchTeam(DB *gorm.DB, team model.Team) {
	// 处理比赛战队
	var count int = 0
	DB.Model(&model.Team{}).Where("team_name = ?", team.TeamName).Count(&count)
	// 存在战队记录就修改，不存在就新建战队记录
	if count == 0 {
		team.TeamBizId = utils.GenerateModuleBizID("TM")
		team.CreatedTime = time.Now()
		team.Insert(DB)
	} else {
		DB.Model(model.Team{}).Where("team_name = ?", team.TeamName).Updates(team)
	}

	// 处理战队相关的队员
	var queryTeam = model.Team{}
	DB.Where("team_name = ?", team.TeamName).Find(&queryTeam)
	//fmt.Println(queryTeam)
	//fmt.Println(queryTeam.TeamBizId)

	if len(team.Players) > 0 {
		for _, player := range team.Players {
			count = 0
			DB.Model(&model.Player{}).Where("nick_name = ?", player.NickName).Count(&count)
			//fmt.Println(player.Name, count)

			if count == 0 {
				player.TeamBizId = team.TeamBizId
				player.PlayerBizId = utils.GenerateModuleBizID("PR")
				player.CreatedTime = time.Now()
				player.Insert(DB)
			}
		}
	}

}
