package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio,omitempty"`
}

func main() {
	userData := []byte(`{"id":1,"username":"Bob","email":"bob@gmail.com"}`)
	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		panic(err)
	}
	fmt.Println(user)

	var userMap map[string]interface{}
	if err := json.Unmarshal(userData, &userMap); err != nil {
		panic(err)
	}
	userID := userMap["id"].(float64)
	fmt.Println(userID)
	fmt.Println(user)

	user = User{ID: 1, Username: "John", Email: "johny@foo.bar"}
	userData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(userData))

	usersData := []byte(`[{"id":1, "username":"Bob", "email":"bob@gmail.com"}, {"id":1, "username":"Bob", "email":"bob@gmail.com"}]`)
	var users []User
	if err = json.Unmarshal(usersData, &users); err != nil {
		panic(err)
	}
	fmt.Println(users)

	type Err struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	type AppError struct {
		Error Err `json:"error"` // 嵌套 Err 结构体，JSON 键名是 error
	}

	// Unmarshalling JSON into struct
	errData := []byte(`{"error":{"code":200, "message":"oops, something went wrong"}}`)
	var appErr AppError
	if err := json.Unmarshal(errData, &appErr); err != nil {
		panic(err)
	}
	fmt.Println(appErr) // 输出：{{200 oops, something went wrong}}

	appErr = AppError{
		Error: Err{
			Code:    200,
			Message: "Some error message",
		},
	}
	errData, err = json.Marshal(appErr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(errData))

}
