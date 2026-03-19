package main

import (
	"fmt"
	"reflect"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("I am %s and %d years old", p.name, p.age)
}

func main() {
	list := make(List, 4) //创建长度为4的List切片（元素初始为nil）
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"Jack", 29}
	list[3] = 1.12

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an Person and its value is %v\n", index, value)
		} else {
			fmt.Println("Unknow type")
		}
	}

	for index, element := range list {
		// type switch:专门用于接口类型的多分支类型判断
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			// reflect.ValueOf(value).Kind()：获取值的底层类型（如 float64）
			fmt.Printf("list[%d] is of a different type %s\n", index, reflect.ValueOf(value).Kind())
		}
	}
}
