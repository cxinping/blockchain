package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/src/model"
	"spider/src/utils"
	"spider/src/utils/parameter"
	"strconv"
	"time"
)

func OperateTournament(dom *goquery.Document) []model.Tournament {
	// 处理赛事数据
	fmt.Println("--- OperateTournament --- ")
	tourResultSet := make([]model.Tournament, 0, 10)
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
			tour.TtName = eventName
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
			tour.TtName = eventName
			tour.TtPic = eventPic
			tourResultSet = append(tourResultSet, tour)
		}

		//fmt.Println("idx=>", idx, ttUrl)
		//fmt.Println("eventName=", eventName)
		//fmt.Println()
	})

	return tourResultSet
}

func OperateUpcomingMatch(dom *goquery.Document) []model.Match {
	// 处理将要比赛的数据
	matchResultSet := make([]model.Match, 0, 10)

	dom.Find(".upcomingMatchesSection").Each(func(idx int, selection *goquery.Selection) {
		match := model.Match{}

		// 比赛时间
		matchDate := selection.Find(".matchDayHeadline").Text()
		//match_date = strings.Replace(match_date, " ", "", -1)
		//match_date_idx := strings.Index(match_date, "-") + 1
		//match_date = string([]rune(match_date)[match_date_idx:len(match_date)])
		fmt.Println("页面部分 idx=>", idx+1, ",matchDate=", matchDate)

		selection.Find(".upcomingMatch").Each(func(i int, selection *goquery.Selection) {
			selDom := selection.Find("div[class*='upcomingMatch']>a").Eq(0)
			matchUrl, _ := selDom.Attr("href")
			matchUrl = parameter.HLTV_INDEX + matchUrl
			match.MatchUrl = matchUrl
			fmt.Println("\tmatch_url=>", matchUrl)

			matchDateUnixStr, _ := selection.Find(".matchInfo .matchTime").Attr("data-unix")
			fmt.Println("\tmatchDateUnixStr=", matchDateUnixStr)

			matchDateUnixInt, _ := strconv.ParseInt(matchDateUnixStr, 10, 64)
			matchDateUnixInt = int64(matchDateUnixInt) / 1000
			matchTime := time.Unix(matchDateUnixInt, 0)
			matchTimeStr := matchTime.Format("2006-01-02 15:04")
			fmt.Println("\tmatchTimeStr=", matchTimeStr)

			team1_name := utils.CompressString(selection.Find("div[class='matchTeam team1']").Text())
			team1_pic, _ := selection.Find("div[class='matchTeam team1']").Find("img").Attr("src")
			team2_name := utils.CompressString(selection.Find("div[class='matchTeam team2']").Text())
			team2_pic, _ := selection.Find("div[class='matchTeam team2']").Find("img").Attr("src")
			match_pic, _ := selection.Find(".matchEvent").Find(".matchEventLogoContainer").Find("img").Attr("src") // 比赛的图片logo
			tt_name := selection.Find("div[class='matchEventName gtSmartphone-only']").Text()
			tt_pic, _ := selection.Find(".matchEventLogoContainer").Find("img").Eq(0).Attr("src")
			mapType := selection.Find(".matchMeta").Text()
			// 查询推荐指数,以黑色星星表示推荐数
			var starNum int8 = 0
			selection.Find("div[class='matchRating']").Find("i").Each(func(i int, matchRatingDomSel *goquery.Selection) {
				starClass, _ := matchRatingDomSel.Attr("class")
				if starClass == "fa fa-star" {
					starNum++
				}
			})

			if team1_name != "" && team2_name != "" {
				fmt.Println("\tteam1_name=", team1_name)
				fmt.Println("\tteam1_pic=", team1_pic)
				fmt.Println("\tteam2_name=", team2_name)
				fmt.Println("\tteam2_pic=", team2_pic)
				fmt.Println("\tmatch_pic=", match_pic)
				fmt.Println("\ttt_name=", tt_name)
				fmt.Println("\ttt_pic=", tt_pic)
				fmt.Println("\tmapType=", mapType)
				fmt.Println("\tstarNum=", starNum)
				fmt.Println("")

				match.MatchTime = matchTime
				match.TtName = tt_name
				match.TtPic = tt_pic
				match.Desc = "section-" + strconv.Itoa(idx+1)
				match.MapType = mapType
				match.SuggestIdx = uint8(starNum)
				match.Team1Name = team1_name
				match.Team2Name = team2_name
				matchResultSet = append(matchResultSet, match)
			}

		})
		fmt.Println("")

	})
	return matchResultSet
}

func OperateLivingMatch(dom *goquery.Document) []model.Match {
	// 处理正在比赛的数据
	liveMatchSectionDom := dom.Find(".liveMatchesSection")
	matchResultSet := make([]model.Match, 0, 10)

	if liveMatchSectionDom != nil {
		// 赛事名称
		match_name := liveMatchSectionDom.Find(".upcoming-headline").Text()
		fmt.Printf("\t正在比赛的赛事名称=%v\n", match_name)

		liveMatchSectionDom.Find(".liveMatch-container").Each(func(idx int, selection *goquery.Selection) {
			match := model.Match{}
			fmt.Printf("\tindex=%d, match=> %T\n", idx+1, match)
			match_url, _ := selection.Find("a[class='match a-reset']").Attr("href")
			match.MatchUrl = parameter.HLTV_INDEX + match_url
			fmt.Println("\t正在比赛的地址=> ", "https://www.hltv.org"+match_url)
			tt_name := selection.Find("div[class='matchEventName gtSmartphone-only']").Text()
			match.TtName = tt_name
			fmt.Println("\t赛事名字=>", tt_name)
			map_type := selection.Find("div[class='matchMeta']").Text()
			match.MapType = map_type
			fmt.Println("\t地图类型=>", map_type)

			selection.Find("div[class='matchTeam']").Each(func(i int, selection *goquery.Selection) {
				team_name := utils.CompressString(selection.Find("div[class='matchTeamName text-ellipsis']").Text())
				if i == 0 {
					team1_name := team_name
					team1_pic, _ := selection.Find("div[class='matchTeamLogoContainer']").Find("img").Attr("src")
					team1_playing_score := selection.Find("div[class='currentMapScore trailing']").Find("span").Text()
					fmt.Println("\tteam1_name=", team1_name)
					fmt.Println("\tteam1_pic=", team1_pic)
					fmt.Println("\tteam1_playing_score=", team1_playing_score)
				} else if i == 1 {
					team2_name := team_name
					team2_pic, _ := selection.Find("div[class='matchTeamLogoContainer']").Find("img").Attr("src")
					team2_playing_score := selection.Find("div[class='currentMapScore trailing']").Find("span").Text()
					fmt.Println("\tteam2_name=", team2_name)
					fmt.Println("\tteam2_pic=", team2_pic)
					fmt.Println("\tteam2_playing_score=", team2_playing_score)
				}
			})

			fmt.Println("")
			matchResultSet = append(matchResultSet, match)
		})
		//fmt.Println()
	}

	return matchResultSet
}
