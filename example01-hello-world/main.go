package main

import "fmt"

func main() {
	a := 1
	fmt.Println(HelloWorld("appleboy"))
	fmt.Println("一天就学会Go语言")
	if a >= 1 {
		fmt.Println("a>=1")
	}
}

func HelloWorld(user_name string) string {
	return fmt.Sprintf("H1 %s", user_name)
}
