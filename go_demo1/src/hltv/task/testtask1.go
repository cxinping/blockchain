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

func add_task(c *cron.Cron) {
	c.AddFunc("*/15 * * * * ?", func() { fmt.Println("Every 15 seconds", time.Now().Format("2006-01-02 15:04:05")) })
}

var CORN = cron.New(cron.WithSeconds())

func InitTask() {
	fmt.Println("--- start task ---", time.Now().Format("2006-01-02 15:04:05"))

	//fmt.Printf("c = %T", c)

	// */5 * * * * ?
	CORN.AddFunc("0 */1 * * * ?", func() {
		fmt.Println("* 执行任务", time.Now().Format("2006-01-02 15:04:05"))
	})

	CORN.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
	})

	CORN.Start()

	//add_task(CORN)

	defer CORN.Stop()
	for {
		time.Sleep(time.Second)
	}
}

type MyJob func()

func (f MyJob) Run() {
	fmt.Println("myJob")
}

func ExecTask1() {
	c := cron.New()
	c.AddJob("@every 1s", MyJob(func() {
		fmt.Println("myjob")
	}))

	c.Start()
	select {}
}
