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

func OperateUpcomingMatch(dom *goquery.Document) []model.Match {
	// 处理将要比赛的数据
	matchResultSet := make([]model.Match, 0, 10)

	// .upcomingMatchesAll .upcomingMatchesSection
	dom.Find(".upcomingMatchesSection").Each(func(idx int, selection *goquery.Selection) {
		match := model.Match{}
		sel_ls := selection.Find("div[class*='upcomingMatch']>a").Eq(0)
		match_url, _ := sel_ls.Attr("href")
		match_url = parameter.HLTV_INDEX + match_url
		match.Match_url = match_url
		fmt.Println("***  match_url=>", match_url, sel_ls)

		// 比赛时间
		match_date := selection.Find(".matchDayHeadline").Text()
		//match_date = strings.Replace(match_date, " ", "", -1)
		//match_date_idx := strings.Index(match_date, "-") + 1
		//match_date = string([]rune(match_date)[match_date_idx:len(match_date)])
		fmt.Println("idx=>", idx+1, ",match_date=", match_date)

		selection.Find(".upcomingMatch").Each(func(i int, selection *goquery.Selection) {
			//match_time := selection.Find(".matchInfo .matchTime").Text()
			//fmt.Println("\tmatch_time=", match_time)

			match_date_unix_str, _ := selection.Find(".matchInfo .matchTime").Attr("data-unix")
			fmt.Println("\tmatch_date_unix_str=", match_date_unix_str)
			match_date_unix_int, _ := strconv.ParseInt(match_date_unix_str, 10, 64)
			match_date_unix_int = int64(match_date_unix_int) / 1000
			match_time := time.Unix(match_date_unix_int, 0)
			match_time_str := match_time.Format("2006-01-02 15:04")
			fmt.Println("\tmatch_time_str=", match_time_str)
			team1_name := utils.CompressString(selection.Find("div[class='matchTeam team1']").Text())
			team1_pic, _ := selection.Find("div[class='matchTeam team1']").Find("img").Attr("src")
			team2_name := utils.CompressString(selection.Find("div[class='matchTeam team2']").Text())
			team2_pic, _ := selection.Find("div[class='matchTeam team2']").Find("img").Attr("src")
			match_pic, _ := selection.Find(".matchEvent").Find(".matchEventLogoContainer").Find("img").Attr("src") // 比赛的图片logo
			tt_name := selection.Find("div[class='matchEventName gtSmartphone-only']").Text()
			fmt.Println("\tteam1_name=", team1_name)
			fmt.Println("\tteam1_pic=", team1_pic)
			fmt.Println("\tteam2_name=", team2_name)
			fmt.Println("\tteam2_pic=", team2_pic)
			fmt.Println("\tmatch_pic=", match_pic)
			fmt.Println("\ttt_name=", tt_name)
			fmt.Println("")

			match.Match_time = match_time
			match.TT_name = tt_name
			match.Desc = strconv.Itoa(idx + 1)
			matchResultSet = append(matchResultSet, match)
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
			match.Match_url = parameter.HLTV_INDEX + match_url
			fmt.Println("\t正在比赛的地址=> ", "https://www.hltv.org"+match_url)
			tt_name := selection.Find("div[class='matchEventName gtSmartphone-only']").Text()
			match.TT_name = tt_name
			fmt.Println("\t赛事名字=>", tt_name)
			map_type := selection.Find("div[class='matchMeta']").Text()
			match.Map_type = map_type
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
