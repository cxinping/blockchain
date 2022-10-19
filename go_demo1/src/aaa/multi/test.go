package multi

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WaitGroupStart(url string) {
	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			Spider(url, i)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("WaitGroupStart Time %s\n ", elapsed)
}

func Spider(url string, idx int) {
	fmt.Println("url=", url, ", idx=", idx)
	time.Sleep(1 * time.Second)
}
