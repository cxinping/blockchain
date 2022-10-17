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

func CrawlMatches() (err error) {
	DB := config.GetDB() // 初始化数据库句柄

	// 爬取赛事信息
	c := colly.NewCollector(
		// 允许重复访问
		colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		//反爬虫，通过随机改变 user-agent,
		r.Headers.Set("User-Agent", utils.RandomString())
		//fmt.Println("OnRequest")
		fmt.Println("访问网页 => ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))

		toursResultSet := OperateTournament(dom) // 处理赛事数据
		operateTournaments(DB, toursResultSet)

		var matchResultSet []model.Match
		matchResultSet = OperateLivingMatch(dom) // 处理正在进行的赛果/赛程的数据
		operateLivingMatches(DB, matchResultSet)

		matchResultSet = OperateUpcomingMatch(dom) // 处理将要进行的赛果/赛程的数据
		operateUpcomingMatches(DB, matchResultSet)

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
	if tts != nil {
		for _, tour := range tts {
			// 多次爬取网页数据时，避免插入重复数据
			var queryTour = model.Tournament{}
			DB.Where("tt_name = ?", tour.TtName).Find(&queryTour)
			if queryTour.TtName != tour.TtName {
				tour.TtBizId = utils.GenerateModuleBizID("TT")
				tour.CreatedTime = time.Now()
				tour.Insert(DB)
			}

		}
	}

}
