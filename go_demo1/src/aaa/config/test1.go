package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func TestConfig1() {
	workDir, _ := os.Getwd()
	fmt.Println("workDir=> ", workDir)
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取YML配置出错", err)
	}
	fmt.Println("db.driver => ", viper.Get("db.driver"))
	fmt.Println("db.host => ", viper.Get("db.host"))

}
