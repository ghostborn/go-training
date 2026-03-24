package main

import "testing"

func TestIsMyError(t *testing.T) {
	err := ErrUserNameExist{UserName: "appleboy"}

	ok := IsErrUserNameExist(err)
	if !ok {
		t.Fatal("testing error")
	}
}

// 在代码所在目录执行命令：go test -v  -v：显示详细测试日志（不加则仅显示是否通过）
