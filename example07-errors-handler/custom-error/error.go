package main

import "fmt"

type ErrUserNameExist struct {
	UserName string
}

// 只要一个类型实现了Error() string方法，就满足error接口，可作为错误返回 / 使用。

func (e ErrUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}

func IsErrUserNameExist(err error) bool {
	_, ok := err.(ErrUserNameExist)
	return ok
}
