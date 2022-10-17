package main

import (
	"fmt"
	"go_demo1/src/aaa/common"
	"go_demo1/src/aaa/http2"
	"go_demo1/src/hltv/logger"
	"go_demo1/src/hltv/spider"
	"go_demo1/src/hltv/task"
	"strconv"
	"strings"
	"time"
)

func test_db() {
	//db.TestCreateTable()
	//db.TestDropTable()

	//db.TestCreateTable2()
	//db.TestDelete1()

	//db.TestSelect1()

	//例子2
	//db2.TestInitTable()
	//db2.TestInitTable2()
	//db2.TestCreateUser()
	//db2.TestUpdateUser()
	//db2.TestDelete()
	//db2.TestSelect()

}

func test1() {
	common.SayHello()

}

func calTime() {
	//计算耗时操作
	start := time.Now()
	time.Sleep(time.Second * 2)

	fmt.Println(time.Now().Sub(start))
}

func test_http() {
	//爬虫
	//http.TestHttp1()
	//http.TestHttp2()
	//http.TestHttp3()
	//http.TestHttp4()
	//http.TestHttp5()

	//http.Example1()
	//http.Example2()
	//http.Example3()

	//http2.Example1()
	//http2.Example2()
	http2.Example3()

}

func test_crawl() {
	spider.CrawlMatches()
}

func test_timeunix() {
	//timeStamp := time.Now().Unix()
	//fmt.Println(timeStamp)
	fmt.Println(1665792000000) // 1665792000000
	timeStr := time.Unix(1665792000000/1000, 0).Format("2006-01-02 15:04")
	fmt.Println("*** timeStr=> ", timeStr)
	//str1 := "1665475200000"  // 2022-10-11 15:30
	//fmt.Println(str1[0 : len(str1)-3])
}

func test_split() {
	time_str := "2022-10-11 15:30"
	time_sep := "-"
	arr := strings.Split(time_str, time_sep)
	fmt.Println(arr)

	//f := func(c rune) bool {
	//	if c == '*' || c == '@' || c == 'f' || c == ' ' || c == '二' {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//s := "@a*b@@c**d## e$f二%ag*"
	//result := strings.FieldsFunc(s, f)
	//fmt.Printf("result:%q", result)

	f := func(c rune) bool {
		if c == ' ' || c == '-' || c == ':' {
			return true
		} else {
			return false
		}
	}
	s := "2022-10-11 15:30"
	results := strings.FieldsFunc(s, f)
	fmt.Printf("results=:%q\n", results)
	fmt.Println(results[0], results[1])
	// 2022年10月11日19:21:30秒执行任务
	fmt.Println("执行的表达式 => " + getCron("*", "10", "11", "19", "21", "10"))
}

func getCron(day_of_week string, month, day string, hour string, minutes string, seconds string) string {
	/**
	Field name   | Mandatory? | Allowed values  | Allowed special characters
	----------   | ---------- | --------------  | --------------------------
	Seconds      | Yes        | 0-59            | * / , -
	Minutes      | Yes        | 0-59            | * / , -
	Hours        | Yes        | 0-23            | * / , -
	Day of month | Yes        | 1-31            | * / , - ?
	Month        | Yes        | 1-12 or JAN-DEC | * / , -
	Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

	*/
	return seconds + " " + minutes + " " + hour + " " + day + " " + month + " " + day_of_week
}

func test_task() {
	task.InitTask()
	//task.ExecTask2()
}

func msToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))
	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

	return tm, nil
}

func testConfig() {
	//config.TestConfig1()
	logger.TestLog()
}

func main() {
	//test1()
	//test_db()

	//test_split()

	//test_timeunix()
	//rst, _ := msToTime("1665844800000")
	//fmt.Printf("%v, %T", rst, rst)

	test_http()
	//test_crawl()
	//test_task()
	//common.TestDelay()

	//testConfig()
}
