package main

import "fmt"

func main() {
	showAllElements(1, 2, 3)

	fmt.Println("~~~~~~~~~~~")

	slice := []int{10, 11, 12}
	showAllElements(slice...)

	fmt.Println("~~~~~~~~~~~")

	slice2 := []int{13, 14, 15}

	newSlice := append(slice, slice2...)
	showAllElements(newSlice...)

	fmt.Println("~~~~~~~~~~~")

	fmt.Println(slice)
	changeValue(slice)
	fmt.Println(slice) //Slice has been changed

	fmt.Println("~~~~~~~~~~~")
	fmt.Println(slice2)
	changeCapacity(slice2)
	fmt.Println(slice2) //Nothing changes

	fmt.Println("~~~~~~~~~~~")
	fmt.Println(slice2)
	brnewSlice := changeCapacity2(slice2)
	fmt.Println(brnewSlice) //That is why we create brand-new slice

}

func showAllElements(nums ...int) {
	for _, value := range nums {
		fmt.Printf("Value = %#v\n", value)
	}
}

func changeValue(slice []int) {
	slice[0] = 88
}

func changeCapacity(slice []int) {
	slice = append(slice, 1, 2)
}

func changeCapacity2(slice []int) []int {
	slice = append(slice, 16, 17)
	return slice
}
