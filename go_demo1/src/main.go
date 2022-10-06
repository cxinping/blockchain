package main

import (
	common "go_demo1/src/aaa/common"
	"go_demo1/src/aaa/db2"
)

func test_db() {
	//db.TestCreateTable()
	//db.TestDropTable()

	//db.TestCreateTable2()
	//db.TestDelete1()

	//db.TestSelect1()

	//db2.TestInit2()
	db2.TestCreateUser()

}

func test1() {
	common.SayHello()
}

func main() {
	//test1()

	test_db()

}
