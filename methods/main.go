package main

import "fmt"

type Square struct {
	side int
}

func main() {
	First := Square{1}
	Second := Square{2}

	First.Perimeter()
	Second.Perimeter()

	pFirst := &First
	pSecond := &Second

	pFirst.Scale(2)
	pSecond.Scale(5)
}

func (s Square) Perimeter() { //Value receiver means that we work only with copy of the object(cannot change entities inside structure)
	fmt.Printf("Side of square = %v\n", s.side)
	fmt.Printf("Perimeter = %v\n", s.side*4)
	fmt.Println("~~~~~~~~~~~~~~~")
}

func (s *Square) Scale(multiplier int) { //Pointer receiver means that we can affect on entities inside structure
	fmt.Printf("Side of square before scaling = %v, muliplier = %v\n", s.side, multiplier)
	s.side *= multiplier
	fmt.Printf("Side of square after scaling = %v\n", s.side)
	fmt.Println("~~~~~~~~~~~~~~~")
}
