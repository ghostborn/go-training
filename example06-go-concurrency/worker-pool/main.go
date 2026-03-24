package main

import (
	"fmt"
	"sync"
	"time"
)

const concurrency = 3

func main() {
	tasks := make(chan int, 100)
	go func() {
		for j := 1; j <= 9; j++ {
			tasks <- j
		}
		close(tasks)
	}()

	results := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(concurrency)
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= concurrency; i++ {
		go func(id int) {
			defer wg.Done() // 协程退出前，通知wg:该worker完成

			for t := range tasks {
				fmt.Println("worder", id, "processing job", t)
				results <- t * 2 // 处理结果写入result通道
				time.Sleep(time.Second)
			}

		}(i)
	}

	for r := range results {
		fmt.Println("result", r)
	}

}
