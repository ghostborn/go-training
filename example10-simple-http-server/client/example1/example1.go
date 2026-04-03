package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://dummyjson.com/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
