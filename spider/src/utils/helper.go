package utils

import (
	"fmt"
	"github.com/go-basic/uuid"
	"math/rand"
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

func CompressString(input_str string) string {
	// 压缩字符串
	input_str = strings.Replace(input_str, "\n", "", -1)
	input_str = strings.Trim(input_str, " ")
	return input_str
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
	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

	return tm, nil
}
