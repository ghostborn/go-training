package main

import "fmt"

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	a = iota
	b
)

func main() {
	// fmt.Println(monday)
	// fmt.Println(tuesday)
	// fmt.Println(wednesday)
	// fmt.Println(thursday)
	// fmt.Println(friday)
	// fmt.Println(saturday)
	// fmt.Println(sunday)
	// fmt.Println(a)
	// fmt.Println(b)

	stringSlice1 := []string{"1", "2", "3", "4"}
	// fmt.Println(stringSlice1[:2])
	// fmt.Println(stringSlice1[2:])
	//

	stringSlice2 := make([]string, 5, 10)
	copy(stringSlice2, stringSlice1)
	fmt.Println(stringSlice2)

	stringSlice2 = append(stringSlice2, "5", "6")
	fmt.Println(stringSlice2)
	fmt.Println(stringSlice2[5])
	fmt.Println(stringSlice2[6])
	stringSlice2 = append(stringSlice2, []string{"7", "8"}...)
	fmt.Println(stringSlice2)
	fmt.Println(stringSlice2[7])
	fmt.Println(stringSlice2[8])

}
