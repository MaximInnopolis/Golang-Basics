package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Builder struct {
	Person
}

func (s Person) Say() {
	fmt.Println("Hello world!")
}

func main() {
	builder := Builder{Person{name: "Max", age: 21}}
	fmt.Printf("Type = %T, value = %v\n", builder, builder)
	fmt.Println(builder)
	builder.Say()
}
