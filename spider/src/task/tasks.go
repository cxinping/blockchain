package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type MyJob func()

func (f MyJob) Run() {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("* myJob 查询比赛列表修改状态 ", currentTime)
}

var Cron = cron.New(cron.WithSeconds())

func ExecTasks() {
	// 定时执行调度任务

	// 查询比赛列表修改状态
	fmt.Println("* 查询比赛列表修改状态")
	Cron.AddJob("@every 5s", MyJob(func() {
		fmt.Println("myjob")
	}))

	Cron.Start()
	select {} //阻塞主线程停止
}
