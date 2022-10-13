package utils

import (
	"math/rand"
	"strings"
)

func RandomString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CompressString(input_str string) string {
	input_str = strings.Replace(input_str, "\n", "", -1)
	input_str = strings.Trim(input_str, " ")
	return input_str
}
