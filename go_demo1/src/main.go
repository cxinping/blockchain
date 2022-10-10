package main

import (
	"fmt"
	common "go_demo1/src/aaa/common"
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

func test_http() {
	//爬虫
	//http.TestHttp1()
	//http.TestHttp2()
	//http.TestHttp3()
	//http.TestHttp4()
	//http.TestHttp5()

}

func calTime() {
	//计算耗时操作
	start := time.Now()
	time.Sleep(time.Second * 2)

	fmt.Println(time.Now().Sub(start))
}

func test_crawl() {
	spider.Crawl_index()
}

func main() {
	//test1()
	//test_db()
	//test_http()

	test_crawl()
}
