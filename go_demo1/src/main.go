package main

import (
	"fmt"
	common "go_demo1/src/aaa/common"
	"go_demo1/src/aaa/http"
	"go_demo1/src/hltv/spider"
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
	http.Example3()

}

func test_crawl() {
	spider.CrawlMatches()

}

func test2() {
	timeStamp := time.Now().Unix()
	fmt.Println(timeStamp)
	fmt.Println(1665475200000) // 1665473400000 1665475200000 1665478800000 1665475200
	timeStr := time.Unix(1665473400, 0).Format("2006-01-02 15:04")
	fmt.Println("*** timeStr=> ", timeStr)
	//str1 := "1665475200000"
	//fmt.Println(str1[0 : len(str1)-3])
}

func main() {
	//test1()
	//test_db()
	//test_http()

	//test2()
	test_crawl()

}
