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
	c.AddFunc("*/30 * * * * ?", func() { fmt.Println("Every 30 seconds", time.Now().Format("2006-01-02 15:04:05")) })
}

var Cron = cron.New(cron.WithSeconds())

func InitTask() {
	fmt.Println("--- start task ---", time.Now().Format("2006-01-02 15:04:05"))

	// */5 * * * * ?
	Cron.AddFunc("13 29 19 11 10 ?", func() {
		fmt.Println("******* 执行任务", time.Now().Format("2006-01-02 15:04:05"))
	})

	Cron.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
	})

	Cron.Start()

	add_task(Cron)

	defer Cron.Stop()
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
	c.AddJob("@every 5s", MyJob(func() {
		fmt.Println("myjob")
	}))

	fmt.Println("111--- main over ---")

	c.Start()
	select {} //阻塞主线程停止

	fmt.Println("222--- main over ---")
}

func ExecTask2() {
	c := cron.New(cron.WithSeconds()) //精确到秒

	//定时任务
	spec := "*/2 * * * * ?" //cron表达式，每秒一次
	c.AddFunc(spec, func() {

		//status := getStatus() //获取定时任务的状态
		//if status == true {
		//	fmt.Println("11111")
		//} else {
		//	c.Stop() //当前定时任务状态已注销
		//}
	})

	c.Start()
	select {} //阻塞主线程停止
}
