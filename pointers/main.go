package main

import "fmt"

func main() {
	var a int = 7
	var pa *int = &a
	fmt.Printf("Type = %T, address = %v, value = %v\n", pa, pa, *pa)
	fmt.Println("~~~~~~~~~")

	var pb *int
	fmt.Printf("Type = %T, address = %v\n", pb, pb)
	fmt.Println("~~~~~~~~~")

	var pointer = new(float32)
	fmt.Printf("Type = %T, address = %v, value = %v\n", pointer, pointer, *pointer)
	fmt.Println("~~~~~~~~~")

	*pointer = 2
	fmt.Printf("Type = %T, address = %v, value = %v\n", pointer, pointer, *pointer)
	fmt.Println("~~~~~~~~~")

	value := 3
	square(&value)
	fmt.Printf("Type = %T, value = %v\n", value, value)
}

func square(value *int) {
	*value *= *value
}
