package main

import "fmt"

func main() {
	multiplier := func(x, y int) int { return x * y }
	fmt.Println(multiplier(1, 2))
	fmt.Println("~~~~~~~~~~~~~~~~~~")

	sum := func(first, second int) int {
		return first + second
	}

	calculate := func(x, y int, action func(a, b int) int) int {
		return action(x, y)
	}

	fmt.Println(calculate(3, 4, sum))
	fmt.Println("~~~~~~~~~~~~~~~~~~")

	divideBy2 := CreateDivider(2)
	fmt.Println(divideBy2(100))
	fmt.Println("~~~~~~~~~~~~~~~~~~")

Label:
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Println("I: ", i, " J: ", j)
			if i >= 4 {
				break Label
			}
		}
	}
}

func CreateDivider(divider int) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider
	}
	return dividerFunc
}
