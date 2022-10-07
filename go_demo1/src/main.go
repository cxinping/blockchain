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

	//例子2
	//db2.TestInitTable()
	//db2.TestInitTable2()
	//db2.TestCreateUser()
	//db2.TestUpdateUser()
	//db2.TestDelete()
	db2.TestSelect()

}

func test1() {
	common.SayHello()
}

func main() {
	//test1()

	test_db()

}
