package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界! 👨‍👩‍👧‍👦"
	for i, r := range s {
		fmt.Printf("位置 %d: 字符 '%c' (Unicode: U+%04X)\n", i, r, r)
	}

	// 获取字符数
	runeCount := utf8.RuneCountInString(s)
	fmt.Printf("字符数: %d\n", runeCount)

	// 获取字节数
	byteCount := len(s)
	fmt.Printf("字节数:%d\n", byteCount)

}
