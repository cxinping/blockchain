package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// 初始化配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err, "读取YML配置出错")
	}
}
