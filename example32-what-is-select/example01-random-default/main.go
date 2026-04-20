package main

import "fmt"

func main() {
	ch := make(chan int, 1) // 创建一个带缓冲的int类型通道

	ch <- 1
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 01")
	default:
		fmt.Println("exit")
	}
}
