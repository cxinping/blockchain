package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	currentTime = time.Now().Format("2006-01-02")
	logFileName = flag.String("log", "./storage/log/"+currentTime+".log", "Log file name")
)

func handle(level string, title string, content interface{}) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	err := os.MkdirAll("./storage/log/", 0766)
	if err != nil {
		fmt.Println(err)
	}

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "Server Failed")
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("[%s] %s : %s \n", level, title, content)
}

func TestLog() {
	//backup_dir := flag.String("b", "/home/default_dir", "backup path")
	//debug_mode := flag.Bool("d", false, "debug mode")
	//flag.Parse()
	//fmt.Println("backup_dir: ", *backup_dir)
	//fmt.Println("debug_mode: ", *debug_mode)

	handle("info", "bbb", "12345")
}
