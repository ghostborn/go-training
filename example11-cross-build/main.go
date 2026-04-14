package main

import (
	"fmt"
	"go-training/example11-cross-build/hello"
	"go-training/example11-cross-build/hello2"
	"go-training/example11-cross-build/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
