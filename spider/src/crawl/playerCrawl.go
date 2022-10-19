package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net"
	"net/http"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"strings"
	"time"
)

// 爬取队友的网页数据
func CrawlPlayerHttp(playerUrl string) {
	resp, err := http.Get(playerUrl)
	if err != nil {
		fmt.Print("http get err", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("http get err", err)
		return
	}

	//fmt.Println(string(body))
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	player := ParseMatchTeamPlayer(dom)
	player.PlayerUrl = playerUrl
	DB := config.GetDB() // 初始化数据库句柄
	OperatePlayer(DB, player)
}

// 爬取队友的网页数据
func CrawlPlayer(playerUrl string) {
	DB := config.GetDB() // 初始化数据库句柄

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
		r.Headers.Set("User-Agent", utils.RandomString())
	})

	c.OnHTML("div.contentCol", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()

		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		player := ParseMatchTeamPlayer(dom)
		player.PlayerUrl = playerUrl
		OperatePlayer(DB, player)
	})

	// 异常处理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		fmt.Println("")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问战队-队员网页 Visited ", r.Request.URL.String())
	})

	c.Visit(playerUrl)
}

func OperatePlayer(DB *gorm.DB, player model.Player) {
	// 处理战队的队员数据
	var playerCount int = 0
	var queryTeam = model.Team{}
	DB.Where("team_name = ?", player.CurrentTeamName).Find(&queryTeam)

	DB.Model(&model.Player{}).Where("nick_name = ?", player.NickName).Count(&playerCount)
	//fmt.Println("CurrentTeamName=", player.CurrentTeamName)
	//fmt.Println("queryTeam.TeamName=", queryTeam.TeamName, queryTeam.TeamBizId)
	//fmt.Printf("queryTeam.TeamBizId=[%v]\n", queryTeam.TeamBizId)

	// 存在队员记录就修改，不存在就新建队员记录
	if playerCount == 0 && queryTeam.TeamBizId != "" {
		//fmt.Println("111 insert player data")
		player.PlayerBizId = utils.GenerateModuleBizID("PR")
		player.TeamBizId = queryTeam.TeamBizId
		player.CreatedTime = time.Now()
		player.Insert(DB)
	} else {
		//fmt.Println("222 update player data")
		DB.Model(model.Player{}).Where("nick_name = ?", player.NickName).Update(player)
	}

}
