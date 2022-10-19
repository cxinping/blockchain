package utils

import (
	"fmt"
	"github.com/go-basic/uuid"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func RandomString() string {
	// 获得任意字符串
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CompressString(inputStr string) string {
	// 压缩字符串
	inputStr = strings.Replace(inputStr, "\n", "", -1)
	inputStr = strings.Trim(inputStr, " ")
	return inputStr
}

func GenerateModuleBizID(name string) string {
	// 创建模块的业务ID
	uuid := uuid.New()
	uuid = strings.Replace(uuid, "-", "", -1)
	return name + "_" + uuid
}

func MsToTime(ms string) (time.Time, error) {
	// 将毫秒值转换为时间
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))
	return tm, nil
}

func Decimal(num float64) float64 {
	// 保留小数点后2位
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func getExecutePath() string {
	return filepath.Dir(os.Args[0])
}
