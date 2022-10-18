package crawl

import (
	"fmt"
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

// 爬取赛事和赛程网页数据
func CrawlTournamentWeb() (err error) {
	DB := config.GetDB() // 初始化数据库句柄

	// 爬取赛事信息
	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
		//fmt.Println("OnRequest")
		//fmt.Println("访问网页 => ", r.URL)
	})

	//爬取赛事和赛果网页数据
	c.OnHTML("div.mainContent", func(e *colly.HTMLElement) {
		content, _ := e.DOM.Html()
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		toursResultSet := OperateTournament(dom) // 处理赛事数据
		operateTournaments(DB, toursResultSet)

		var livingMatchResultSet []model.Match
		livingMatchResultSet = OperateLivingMatch(dom) // 处理正在进行的赛果/赛程的数据
		operateLivingMatches(DB, livingMatchResultSet)

		var upComingMatchResultSet []model.Match
		upComingMatchResultSet = OperateUpcomingMatch(dom) // 处理将要进行的赛果/赛程的数据
		operateUpcomingMatches(DB, upComingMatchResultSet)

		//爬取正在进行的比赛战队和队员信息
		for _, match := range livingMatchResultSet {
			CrawlMatcheWeb(match.MatchUrl)
		}

		//爬取将要进行的比赛战队和队员信息
		for _, match := range upComingMatchResultSet {
			CrawlMatcheWeb(match.MatchUrl)
		}

	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问赛事和赛程网页 Visited ", r.Request.URL.String())

		//bodyData := string(r.Body)
		//dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))
		//toursResultSet := OperateTournament(dom) // 处理赛事数据
		//operateTournaments(DB, toursResultSet)
		//
		//var matchResultSet []model.Match
		//matchResultSet = OperateLivingMatch(dom) // 处理正在进行的赛果/赛程的数据
		//operateLivingMatches(DB, matchResultSet)

		//matchResultSet = OperateUpcomingMatch(dom) // 处理将要进行的赛果/赛程的数据
		//operateUpcomingMatches(DB, matchResultSet)
	})

	err = c.Visit(parameter.TT_MATCH_URL)
	return err
}

func operateLivingMatches(DB *gorm.DB, matches []model.Match) {
	// 批量保存正在比赛的Match
	if matches != nil {
		var count int

		for _, match := range matches {
			// 多次爬取网页数据时，避免插入重复数据
			DB.Model(&model.Match{}).Where("tt_name = ? AND team1_name = ? AND team2_name = ?", match.TtName, match.Team1Name, match.Team2Name).Count(&count)

			if count == 0 {
				match.MatchBizId = utils.GenerateModuleBizID("MH")
				match.CreatedTime = time.Now()
				match.Status = parameter.MATCH_STATUS_LIVE
				var queryTour = model.Tournament{}
				DB.Where("tt_name = ?", match.TtName).Find(&queryTour)
				match.TtBizId = queryTour.TtBizId
				//fmt.Println(idx, match.TT_name)
				match.Insert(DB)
			}
		}
	}
}

func operateUpcomingMatches(DB *gorm.DB, matches []model.Match) {
	// 批量保存将要进行的比赛的Match
	if matches != nil {
		var count int

		for _, match := range matches {
			// 多次爬取网页数据时，避免插入重复数据
			DB.Model(&model.Match{}).Where("tt_name = ? AND team1_name = ? AND team2_name = ?", match.TtName, match.Team1Name, match.Team2Name).Count(&count)

			if count == 0 {
				match.MatchBizId = utils.GenerateModuleBizID("MH")
				match.CreatedTime = time.Now()
				match.Status = parameter.MATCH_STATUS_UNSTARTED
				var queryTour = model.Tournament{}
				DB.Where("tt_name = ?", match.TtName).Find(&queryTour)
				match.TtBizId = queryTour.TtBizId
				//fmt.Println(idx, match.TT_name)
				match.Insert(DB)
			}
		}
	}
}

func operateTournaments(DB *gorm.DB, tts []model.Tournament) {
	// 批量保存赛事Tournament
	if len(tts) > 0 {
		for _, tour := range tts {
			// 多次爬取网页数据时，避免插入重复数据
			var count int = 0
			DB.Model(&model.Tournament{}).Where("tt_name = ?", tour.TtName).Count(&count)

			if count == 0 {
				tour.TtBizId = utils.GenerateModuleBizID("TT")
				tour.CreatedTime = time.Now()
				tour.Insert(DB)
			}

		}
	}
}
