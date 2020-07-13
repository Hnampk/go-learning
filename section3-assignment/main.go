package main

import "fmt"

func main() {
	listInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, item := range listInt {
		result := "even"

		if item%2 != 0 {
			result = "odd"
		}

		fmt.Println(item, "is", result)

	}
}
