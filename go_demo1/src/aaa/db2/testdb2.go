package db2

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func TestInit1() {
	username := "root"      //账号
	password := "123456"    //密码
	host := "192.168.11.12" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "flink"       //数据库名
	timeout := "10s"        //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("* 连接数据库成功")
	}
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	fmt.Println(db)

}

func TestInit2() {
	fmt.Println("--- TestInit2 ---")

	username := "root"      //账号
	password := "123456"    //密码
	host := "192.168.11.12" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "flink"       //数据库名
	timeout := "10s"        //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	sqlDB, err1 := sql.Open("mysql", dsn)
	db, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	//设置全局表名禁用复数
	//db.SingularTable(true)

	fmt.Println(db, err1, err2)

}

//插入数据
func (user *User) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(user)
}

func TestCreateUser() {
	user := User{
		Name: "wangwu",
		Age:  21,
	}
	fmt.Println(user, )
}