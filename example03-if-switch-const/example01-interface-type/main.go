package main

var (
	i interface{}
) //定义全局变量，类型为空接口

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is integer", t)

	case string:
		println("i is string", t)

	case float64:
		println("i is float64", t)

	default:
		println("type not found")

	}
}

func main() {
	i = 100
	convert(i)
	i = float64(66.66)
	convert(i)
	i = "foo"
	convert(i)
	convert(float32(10.0))
}
