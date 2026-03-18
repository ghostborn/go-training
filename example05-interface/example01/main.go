package main

import (
	"fmt"
	"go-training/example05-interface/example01/lexus"
	"go-training/example05-interface/example01/toyota"
)

type car interface {
	Name() string
	Price() float64
	Discount() float64
	Color() string
}

func detail(c car) {
	fmt.Println("================")
	fmt.Println("Name", c.Name())
	fmt.Println("Price", c.Price())
	fmt.Println("Discount", c.Discount())
	fmt.Println("Color", c.Color())
	fmt.Println("================")

}

func main() {
	car1 := toyota.NewCar("car1", 3000, 0.8, "white")
	car2 := lexus.NewCar("car2", 4000, 0.9, "white")
	detail(car1)
	detail(car2)
}
