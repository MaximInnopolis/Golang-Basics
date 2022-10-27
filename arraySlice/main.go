package main

import (
	"fmt"
)

func main() {
	arr := [...]string{"First", "Second", "Third"}

	fmt.Printf("Type = %T, value = %#v\n", arr, arr)
	fmt.Printf("The length of array is %#v\n", len(arr))

	for indx, value := range arr {
		fmt.Printf("Index = %v, Value = %#v\n", indx, value)
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")

	var slice []int
	fmt.Printf("Type = %T, value = %#v\n", slice, slice)

	slicebyMake := make([]int, 0, 2)
	fmt.Printf("Type = %T, value = %#v, length = %#v, capacity = %#v\n",
		slicebyMake, slicebyMake, len(slicebyMake), cap(slicebyMake))

	slicebyMake = append(slicebyMake, 1, 2, 3, 4)
	fmt.Printf("Type = %T, value = %#v, length = %#v, capacity = %#v\n",
		slicebyMake, slicebyMake, len(slicebyMake), cap(slicebyMake))
}
