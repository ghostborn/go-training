package main

import (
	"fmt"
	"sync"
)

func main() {
	foo := func() string {
		return "Hello World1"
	}

	fmt.Println(foo())

	bar := func() {
		fmt.Println("Hello World2")
	}
	bar()

	func() {
		fmt.Println("Hello World3")
	}()

	var wg sync.WaitGroup
	wg.Add(1) // 标记需要等待一个goroutine

	go func(i, j int) {
		defer wg.Done() // 执行完标记完成
		fmt.Println(i + j)
	}(1, 2)
	wg.Wait() // 等待goroutine执行完毕
	// time.Sleep(1000)

}
