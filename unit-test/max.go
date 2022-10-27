package main

import "fmt"

func main() {
	fmt.Println(Max([]int{2, 1, 6, 3, 8, 32, 16}))
}

func Max(array []int) int {
	var max int

	for _, value := range array {
		if value > max {
			max = value
		}
	}

	return max
}
