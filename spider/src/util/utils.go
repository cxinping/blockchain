package util

import (
	"math/rand"
	"strings"
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
