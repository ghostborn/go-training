package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch := make(chan int)
	select { // select的核心规则是：优先执行最先就绪的case分支，哪个通道先能执行读写操作，就走哪个分支
	case <-ch: // 分支1： 监听ch通道的读事件
	case <-timeout: // 分支2: 监听timeout通道的读事件
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 3): // 1s后就绪的超时通道
		fmt.Println("timeout 02")
	}
}
