package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getHTTPResponse() (*response, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, fmt.Errorf("error in http call")
	}
	defer resp.Body.Close()
	// 读取响应体的字节数据
	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error in reading response: %w", err)
	}
	// 初始化结构体指针，将JSON字节数据反序列化为结构体
	structResp := &response{} //创建空的response指针
	err = json.Unmarshal(byteResp, structResp)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshalling response: %w", err)
	}
	return structResp, nil

}

func main() {
	res, err := getHTTPResponse()
	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Printf("res %v", res)
	}
}
