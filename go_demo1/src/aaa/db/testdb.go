package aaa

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

// 默认表名是 Model名称的小写+复数
type Profile struct {
	gorm.Model
	Refer int
	Name  string
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestCreateTable2() {
	//创建表，缺少的列和索引，不会改变现有列的类型或删除列
	//配置MySQL连接参数
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
	//延时关闭数据库连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 迁移 schema
	//db.AutoMigrate(&Product{})

	//for i := 0; i < 1; i++ {
	//	// Create
	//
	//	Code := fmt.Sprintf("%s%d", "a-", i)
	//	prod := &Product{Code: Code, Price: uint(100 + i)}
	//	db.Create(prod)
	//	fmt.Println(prod.ID, prod.Code)
	//}
	//
	//fmt.Println("插入数据成功")

	var prods = []Product{{Code: "jinzhu1", Price: 100}, {Code: "jinzhu2", Price: 101}, {Code: "jinzhu3", Price: 102}}
	db.Create(&prods)

	for _, prod := range prods {

		fmt.Println(prod.ID) // 1,2,3)
	}

	fmt.Println("批量插入数据成功")
}

func TestDelete1() {
	//创建表，缺少的列和索引，不会改变现有列的类型或删除列
	//配置MySQL连接参数
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
	//延时关闭数据库连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var product Product
	fmt.Printf("%T,%v\n", product, product)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)

	//db.Delete(&Product{Code: "D42", Price: 100 }, 36)

	fmt.Println("删除记录成功")
}

func TestSelect1() {
	//创建表，缺少的列和索引，不会改变现有列的类型或删除列
	//配置MySQL连接参数
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
	//延时关闭数据库连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	product := Product{}
	//result := db.Debug().Where("price = ?", 101).First(&product)

	result := db.First(&product)

	fmt.Println(result)
	fmt.Println(product)
}

func TestCreateTable() {
	//创建表，缺少的列和索引，不会改变现有列的类型或删除列
	//配置MySQL连接参数
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
		fmt.Println("* connected successfully")
	}
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//延时关闭数据库连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var result bool
	result = db.Migrator().HasTable(&User{})
	//result  = db.Migrator().HasTable("users")

	fmt.Println(result)

	db.AutoMigrate(&User{})

}

func TestDropTable() {
	//创建表，缺少的列和索引，不会改变现有列的类型或删除列
	//配置MySQL连接参数
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
		fmt.Println("* connected successfully")
	}
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//延时关闭数据库连接
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 如果存在表则删除（删除时会忽略、删除外键约束)
	//db.Migrator().DropTable(&User{})
	//db.Migrator().DropTable("users")

	// 重命名表
	//db.Migrator().RenameTable(&User{}, &UserInfo{})
	//db.Migrator().RenameTable("users", "user_infos")

}
