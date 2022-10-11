package common

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("SayHello 111")
}

func test12(ch chan int) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-ch:
		if timer.Stop() {
			fmt.Println("关闭定时器")
		}
		return true
	default:
		fmt.Println("继续执行定时器")
		return true
	}
}

func TestDelay() {
	ch := make(chan int, 1)
	ch <- 1

	go test12(ch)

	for {

	}

	fmt.Println("--- main over ---")
}
