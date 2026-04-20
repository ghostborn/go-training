package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	select {
	case ch <- 2: // select 不仅能监听通道读取，也能监听通道写入操作的就绪状态；
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
		// 带缓冲通道的写入是否阻塞，取决于「当前已用缓冲 vs 总容量」；
	default:
		fmt.Println("channel blocking")
	}
}
