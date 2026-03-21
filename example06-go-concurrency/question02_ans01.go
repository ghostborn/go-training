package main

// import "fmt"

// func cacl(i int, c chan bool) {
// 	t := 0
// 	for i := 1; i < 100000; i++ {
// 		t++
// 	}
// 	fmt.Println(i, t)
// 	c <- true
// }

// func main() {
// 	c := make(chan bool)
// 	for i := range 10 {
// 		go cacl(i, c)
// 	}

// 	for i := 0; i < 10; i++ {
// 		<-c
// 	}
// 	fmt.Println("End!!!")
// }
