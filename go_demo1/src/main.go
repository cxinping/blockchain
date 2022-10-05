package main

import (
	common "go_demo1/src/aaa/common"
	db "go_demo1/src/aaa/db"
)

func test_db() {
	//db.TestCreateTable()
	//db.TestDropTable()

	db.TestCreateTable2()
}

func test1() {
	common.SayHello()
}

func main() {
	//test1()

	test_db()

}
