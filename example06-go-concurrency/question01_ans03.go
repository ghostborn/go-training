package main

// import "fmt"

// func main() {
// 	c := make(chan bool) // 创建无缓冲的bool类型通道
// 	go func() {
// 		fmt.Println("let 's go")
// 		c <- true
// 		close(c)
// 	}()

// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }
