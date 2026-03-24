package main

import (
	"fmt"
	"time"
)

type logEntry struct {
	message string
	time    time.Time
}

func main() {
	// 创建带缓冲的日志通道（100）
	logCh := make(chan logEntry, 100)
	// 创建无缓冲的退出信号通道（空结构体struct{}：仅作信号，无数据意义）
	doneCh := make(chan struct{})
	go logger(logCh, doneCh)
	// 向日志通道发送两条日志
	logCh <- logEntry{"App Start", time.Now()}
	logCh <- logEntry{"App End", time.Now()}

	// 短暂休眠，确保日志协程有时间处理完日志（避免主协程过快发退出信号）
	time.Sleep(100 * time.Millisecond)

	// 发送退出信号（空结构体值，仅触发通道可读）
	doneCh <- struct{}{}

}

func logger(logCh <-chan logEntry, doneCh <-chan struct{}) {
	// 定义标签Loop，用于跳出外层for循环
Loop:
	for {
		select {
		case entry := <-logCh:
			// case1：logCh有日志可读 → 打印格式化日志
			fmt.Printf("%v:%v\n", entry.time.Format("2006-01-02T01:01:01"), entry.message)
		case <-doneCh:
			fmt.Println("break the select loop")
			break Loop
		}
	}
	fmt.Println("exit the logger")
}
