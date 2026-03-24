package main

import (
	"fmt"
	"strconv"
)

func main() {
	var test = map[string]interface{}{
		"test01": []interface{}{"a", "b"},
		"test02": []int{1, 2},
	}

	if slice1, ok := test["test01"].([]interface{}); ok {
		oids1 := make([]string, len(slice1))
		for i, v := range slice1 {
			if s, ok := v.(string); ok {
				oids1[i] = s
			}
		}
		fmt.Println(oids1)
	}

	if slice2, ok := test["test02"].([]int); ok {
		oids2 := make([]string, len(slice2))
		oids3 := make([]int, len(slice2))

		// for i, v : = range slice2{
		// 	oids2[i] = strconv.Itoa(v)
		// 	oids3[i] = v
		// }

		for i, v := range slice2 {
			oids2[i] = strconv.Itoa(v)
			oids3[i] = v
		}

		fmt.Println(oids2)
		fmt.Println(oids3)
	}

}
