package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T, value: %#v\n", arr, arr)

	fmt.Println("~~~~~~~~~~~~~~~~~")

	slice := arr[1:3] //From array to slice
	fmt.Printf("Type: %T, value: %#v\n", slice, slice)
	fmt.Printf("Length: %#v, Capacity: %#v\n", len(slice), cap(slice))

	fmt.Println("~~~~~~~~~~~~~~~~~")

	fullSlice := arr[:] //arr[0:5]
	fmt.Printf("Type: %T, value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Length: %#v, Capacity: %#v\n", len(fullSlice), cap(fullSlice))

	fmt.Println("~~~~~~~~~~~~~~~~~")

	sliceFromSLice := fullSlice[:3]
	fmt.Printf("Type: %T, value: %#v\n", sliceFromSLice, sliceFromSLice)
	fmt.Printf("Length: %#v, Capacity: %#v\n", len(sliceFromSLice), cap(sliceFromSLice))

	fmt.Println("~~~~~~~~~~~~~~~~~")

	arr[1] = 500
	fmt.Printf("Type: %T, value: %#v\n", slice, slice)
	fmt.Printf("Type: %T, value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Type: %T, value: %#v\n", sliceFromSLice, sliceFromSLice)

	fmt.Println("~~~~~~~~~~~~~~~~~")
	//To copy elements from one slice to another:

	source := []string{"First", "Second", "Third"}
	rightCopy := append(make([]string, 0, len(source)), source...)
	fmt.Printf("Type: %T, value: %#v\n", rightCopy, rightCopy)
	fmt.Printf("Length: %#v, Capacity: %#v\n", len(rightCopy), cap(rightCopy))

	fmt.Println("~~~~~~~~~~~~~~~~~")
	//To delete elements from slice :

	source = append(source[:1], source[2:]...)
	fmt.Printf("Type: %T, value: %#v\n", source, source)
	fmt.Printf("Length: %#v, Capacity: %#v\n", len(source), cap(source))
}
