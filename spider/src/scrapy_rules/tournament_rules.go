package scrapy_rules

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

func SetTournamentCallback(getTournamentC *colly.Collector) {
	// 爬取赛事的网页数据
	DB := config.GetDB() // 初始化数据库句柄

	getTournamentC.OnResponse(func(r *colly.Response) {
		fmt.Println("访问赛事和赛程网页 Visited ", r.Request.URL.String())
		bodyData := string(r.Body)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(bodyData))
		toursResultSet := OperateTournament(dom) // 处理赛事数据
		operateTournaments(DB, toursResultSet)
	})

}

func OperateTournament(dom *goquery.Document) []model.Tournament {
	// 处理赛事数据
	//fmt.Println("--- OperateTournament --- ")
	tourResultSet := make([]model.Tournament, 0)
	//fmt.Println(tourResultSet, dom)

	eventDom := dom.Find("div[class='events-container']")
	eventDom.Find("a[class='filter-button-link']").Each(func(idx int, selection *goquery.Selection) {
		tour := model.Tournament{}
		ttUrl, _ := selection.Attr("href")
		ttUrl = parameter.HLTV_INDEX + ttUrl
		eventName := selection.Find(".featured-event-tooltip-content").Text()
		eventPicDom := selection.Find("div[class='event-button  tooltip-parent']").Find("img").Eq(0)
		eventPic, _ := eventPicDom.Attr("src")

		if eventName != "" {
			tour.TtUrl = ttUrl
			tour.TtName = utils.CompressString(eventName)
			tour.TtPic = eventPic
			tourResultSet = append(tourResultSet, tour)
		}
		//fmt.Println("\tidx=>", idx, ", ttUrl=", ttUrl, parameter.HLTV_INDEX+ttUrl)
		//fmt.Println("\teventName=", eventName)
	})

	eventDom.Find("div[class='event-filter-popup']").Find("a").Each(func(idx int, selection *goquery.Selection) {
		ttUrl, _ := selection.Attr("href")
		ttUrl = parameter.HLTV_INDEX + ttUrl
		eventName := selection.Find(".event-name").Text()
		eventPic, _ := selection.Find(".event-img").Find("img").Eq(0).Attr("src")
		tour := model.Tournament{}

		if eventName != "" {
			tour.TtUrl = ttUrl
			tour.TtName = utils.CompressString(eventName)
			tour.TtPic = eventPic
			tourResultSet = append(tourResultSet, tour)
		}

		//fmt.Println("idx=>", idx, ttUrl)
		//fmt.Println("eventName=", eventName)
		//fmt.Println()
	})

	return tourResultSet
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
