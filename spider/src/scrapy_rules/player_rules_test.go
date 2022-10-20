package scrapy_rules

import (
	"fmt"
	"runtime"
	"spider/src/config"
	"testing"
	"time"
)

func init() {
	fmt.Printf("本台电脑是 %d 核的CPU\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	config.InitDB()
}

func TestSetPlayerCallback(t *testing.T) {
	start := time.Now()
	playerUrls := make([]string, 0)
	playerUrls = append(playerUrls, "https://www.hltv.org/player/11205/stadodo")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20463/ddias")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20465/arrozdoce")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/20743/suka")
	playerUrls = append(playerUrls, "https://www.hltv.org/player/21014/ag1l")

	for _, playerUrl := range playerUrls {
		ScrapyPlayerInfomation(playerUrl)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("该函数执行完成耗时：", delta)
}
