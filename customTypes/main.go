package main

import (
	"fmt"
	"time"
)

type OurString string
type OurInt int

type Person struct {
	Name string
	Age  int
}

func main() {
	var customString OurString

	fmt.Printf("Type = %T, value = %v\n", customString, customString)
	fmt.Println("~~~~~~~~~~~~~~~")

	customInt := OurInt(1)
	fmt.Printf("Type = %T, value = %v\n", customInt, customInt)
	fmt.Printf("Type = %T, value = %v\n", int(customInt), int(customInt))
	fmt.Println("~~~~~~~~~~~~~~~")

	unnamedStruct := struct {
		name, surname, date string
	}{
		name:    "Max",
		surname: "Smith",
		date:    fmt.Sprintf("%s", time.Now()),
		//date: time.Now().String(),
	}
	fmt.Println(unnamedStruct)
}
