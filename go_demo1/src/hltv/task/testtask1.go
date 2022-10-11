package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func current_time() {
	current_time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("当前时间", current_time)
}

func TestTask1() {
	fmt.Println("--- start task ---")

	c := cron.New(cron.WithSeconds())
	// */5 * * * * ?
	c.AddFunc("1 * * * *", func() {

		fmt.Println("* 执行任务", time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
	defer c.Stop()
	for {
		time.Sleep(time.Second)
	}
}
