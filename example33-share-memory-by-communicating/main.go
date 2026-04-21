package main

import (
	"fmt"
	"sync"
)

func addByShareMemory(n int) []int {
	var ints []int        // 共享切片（所有goroutine都会修改）
	var wg sync.WaitGroup //等待组：等待所有goroutine执行完毕
	var mux sync.Mutex    // 互斥锁：保护共享切片的并发修改

	wg.Add(n) // 注册n个待完成的goroutine

	for i := 0; i < n; i++ {
		go func(i int) { // 启动第i个goroutine
			defer wg.Done()        // 函数退出时，标记该goroutine完成
			mux.Lock()             // 加锁：禁止其他goroutine修改共享切片
			ints = append(ints, i) // 修改共享切片（临界区）
			mux.Unlock()           // 解锁：允许其他goroutine操作
		}(i) // 注意：此处显式传i，避免goroutine捕获循环变量的“延迟绑定”问题
	}
	wg.Wait()   // 阻塞，直到所有wg.Done()被调用（所有goroutine完成）
	return ints // 返回最终收集的切片
}

// 通信（通道）代替共享内存
// 这是 Go 推荐的并发范式：“Do not communicate by sharing memory; instead, share memory by communicating.”（不要通过共享内存通信，而是通过通信共享内存）。
func addByShareCommunicate(n int) []int {
	var ints []int

	channel := make(chan int, n) // 创建带缓冲通道，容量n（避免goroutine阻塞）

	// 启动n个goroutine，向通道写入数据
	for i := 0; i < n; i++ {
		go func(channel chan<- int, order int) {
			channel <- order //向通道写入数字（无锁，通道天然支持并发安全）
		}(channel, i)
	}

	// 从通道读取所有数据，收集到切片
	for i := range channel {
		ints = append(ints, i)
		// 读取n个数据后退出循环（避免通道关闭前阻塞）
		if len(ints) == n {
			break
		}
	}
	close(channel) // 关闭通道（可选，此处通道后续无使用）
	return ints
}

func main() {
	foo := addByShareMemory(10)
	fmt.Println(len(foo))
	fmt.Println(foo)

	foo = addByShareCommunicate(10)
	fmt.Println(len(foo))
	fmt.Println(foo)
}
