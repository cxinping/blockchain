package db2

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func TestInitTable1() *gorm.DB {
	databaseType := "mysql"
	username := "root"      //账号
	password := "123456"    //密码
	host := "192.168.11.12" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "flink"       //数据库名
	timeout := "10s"        //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	//使用gorm链接数据库
	db, err := gorm.Open(databaseType, dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err) //数据库链接失败是致命的错误，链接失败后可以关闭程序了，所以使用logging.Fatal方法
	}

	//设置全局表名禁用复数
	db.SingularTable(true)

	db.AutoMigrate(&User{})
	return db
}

func TestInitTable2() {
	databaseType := "mysql"
	username := "root"      //账号
	password := "123456"    //密码
	host := "192.168.11.12" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "flink"       //数据库名
	timeout := "10s"        //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	//使用gorm链接数据库
	db, err := gorm.Open(databaseType, dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err) //数据库链接失败是致命的错误，链接失败后可以关闭程序了，所以使用logging.Fatal方法
	}

	//设置全局表名禁用复数
	db.SingularTable(true)
	db.LogMode(true) //打印日志，本地调试的时候可以打开看执行的sql语句

	db.DB().SetMaxIdleConns(10)  //设置空闲时的最大连接数
	db.DB().SetMaxOpenConns(100) //设置数据库的最大打开连接数
}

//插入数据
func (user *User) Insert() {
	db = TestInitTable1()
	//fmt.Println(db)

	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	//db.Table("user").Create(user)
	db.Table("user").Debug().Create(user)
}

func TestCreateUser() {
	//新增数据
	for i := 0; i < 10; i++ {
		user := User{
			Name:     fmt.Sprintf("%s%d", "wang", i),
			Age:      31 + i,
			Birthday: time.Now(),
		}
		//fmt.Println(user)
		user.Insert()
	}
}

func TestUpdateUser() {
	db = TestInitTable1()

	//user := User{Name: "xiaoming", Age: 18, Birthday: time.Now()}
	//user.ID = 1
	//db.Model(&user).Debug().Update(user)

	//注意到上面Update中使用了一个Struct，你也可以使用map对象。
	//需要注意的是：使用Struct的时候，只会更新Struct中这些非空的字段。
	//对于string类型字段的""，int类型字段0，bool类型字段的false都被认为是空白值，不会去更新表

	//下面这个更新操作只使用了where条件没有在Model中指定id
	//update user set name='xiaohong' wehre sex=1
	//db.Model(&User{}).Where("id = ?", 2).Update("name", "wangwu").Update("age", 30)

	//var user User
	//user.ID = 3
	//db.Model(&user).Select("name").Update(map[string]interface{}{"name": "lisi", "age": 30})

	user := User{Name: "wangwu", Age: 20}
	user.ID = 4
	db.Model(&user).Omit("name").Update(&user)
}

func TestDelete() {
	fmt.Println("* TestDelete ")
	db = TestInitTable1()
	//delete from user where id=1;
	//var user User
	//user.ID = 1
	//db.Delete(&user)

	//db.Delete(&User{}, 2)

	//批量删除
	db.Where("name LIKE ?", "%lisi%").Delete(User{})

	db.Exec("DELETE FROM user WHERE id <= 5")
}

func TestSelect() {
	db = TestInitTable1()

	// 查询全部数据
	//users := []User{}
	//db.Debug().Find(&users)
	//
	//for idx, user := range users {
	//	fmt.Println(idx, user.ID)
	//}

	users := []User{}
	//指定查询字段
	db.Debug().Select("name,age").Where(map[string]interface{}{"age": 36, "name": "wang5"}).Find(&users)
	for idx, user := range users {
		fmt.Println(idx, user.ID, user.Name, user.Age)
	}
	//fmt.Println(users)

}
