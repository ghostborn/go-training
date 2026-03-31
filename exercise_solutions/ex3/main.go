package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"JIA", "PENG", 32}
	emp2 := Employee{
		firstName: "joe",
		lastName:  "byeden",
		id:        86,
	}
	var emp3 Employee
	emp3.firstName = "mike"
	emp3.lastName = "kobe"
	emp3.id = 42

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
