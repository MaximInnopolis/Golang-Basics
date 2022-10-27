package main

import "fmt"

func main() {
	Greet()
	PersonGreet("Max")
	FioGreet("Ursa", "From Dota")
	result := Sum(1, 2)
	fmt.Println(result)

	sum, mult := SumAndMult(1, 2)
	fmt.Printf("Sum = %v, Mult = %v\n", sum, mult)

	_, nmult := NamedSumAndMult(2, 5)
	fmt.Printf("Mult = %v\n", nmult)
}

func Greet() {
	fmt.Println("Hello World!")
}

func PersonGreet(name string) {
	fmt.Printf("Hello %s\n", name)
}

func FioGreet(name, surname string) {
	fmt.Printf("Hello %s %s\n", name, surname)
}

func Sum(fV, sV int) int {
	return fV + sV
}

func SumAndMult(fV, sV int) (int, int) {
	return fV + sV, fV * sV
}

func NamedSumAndMult(fV, sV int) (sum int64, mult int64) {
	sum = int64(fV + sV)
	mult = int64(fV) * int64(sV)
	return
}
