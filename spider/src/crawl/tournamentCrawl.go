package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"spider/src/config"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strings"
	"time"
)

func CrawlMatches() (err error) {
	// 爬取赛事信息
	base_url := parameter.TT_MATCH_URL // "https://www.hltv.org/matches"
	//fmt.Println("*** 开始爬取hltv的赛事列表 ", base_url)

	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
		//fmt.Println("OnRequest")
		//fmt.Println("url => ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))
		// 初始化数据库句柄
		DB := config.GetDB()

		toursResultSet := OperateTournament(dom) // 处理赛事数据
		saveTournaments(DB, toursResultSet)

		var matchResultSet []model.Match
		matchResultSet = OperateLivingMatch(dom) // 处理正在进行的赛果/赛程的数据
		saveLivingMatches(DB, matchResultSet)

		matchResultSet = OperateUpcomingMatch(dom) // 处理将要进行的赛果/赛程的数据
		saveUpcomingMatches(DB, matchResultSet)

	})

	err = c.Visit(base_url)
	return err
}

func saveLivingMatches(DB *gorm.DB, matches []model.Match) {
	// 批量保存正在比赛的Match
	if matches != nil {
		for _, match := range matches {
			match.MatchBizId = utils.GenerateModuleBizID("MH")
			match.CreatedTime = time.Now()
			match.Status = parameter.MATCH_STATUS_LIVE
			//fmt.Println(idx, match)
			match.Insert(DB)
		}
	}
}

func saveUpcomingMatches(DB *gorm.DB, matches []model.Match) {
	// 批量保存将要进行的比赛的Match
	if matches != nil {
		for _, match := range matches {
			match.MatchBizId = utils.GenerateModuleBizID("MH")
			match.CreatedTime = time.Now()
			match.Status = parameter.MATCH_STATUS_NOT_STARTED
			//fmt.Println(idx, match.TT_name)
			match.Insert(DB)
		}
	}
}

func saveTournaments(DB *gorm.DB, tts []model.Tournament) {
	// 批量保存赛事Tournament
	if tts != nil {
		for _, tour := range tts {
			tour.TtBizId = utils.GenerateModuleBizID("TT")
			tour.CreatedTime = time.Now()
			tour.Insert(DB)
		}
	}
}
