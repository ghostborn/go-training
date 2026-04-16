package main

import (
	"testing"
)

// 基准测试：fmt.Sprintf 方式
func BenchmarkPrintInt2String01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String01(10)
	}
}

// 基准测试：strconv.FormatInt 方式
func BenchmarkPrintInt2String02(b *testing.B) {
	num := int64(10)
	for i := 0; i < b.N; i++ {
		printInt2String02(num)
	}
}

// 基准测试：strconv.Itoa 方式
func BenchmarkPrintInt2String03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String03(10)
	}
}
