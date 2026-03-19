package main

import "fmt"

type IPV4 []byte  //(字节切片) 定义了新的自定义类型

func (s IPV4) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", s[0], s[1], s[2], s[3])
}

func (s IPV4) GoString() string {
	return fmt.Sprintf("%v.%v.%v.%v", s[0], s[1], s[2], s[3])
}

func main() {
	ipv4 := map[string]IPV4{
		"localhost": {127, 0, 0, 1},
		"Google":    {8, 8, 8, 8},
	}

	for i, v := range ipv4 {
		fmt.Printf("name: %s, ip:%v\n", i, v)
		fmt.Printf("debug:%#v\n", v)
	}
}
