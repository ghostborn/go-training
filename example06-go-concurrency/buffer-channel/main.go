package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 2)

	// receive channel
	go func(ch <-chan int) { //只读通道
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	// send channel
	go func(ch chan<- int) { //只写通道
		ch <- 100
		ch <- 101
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
