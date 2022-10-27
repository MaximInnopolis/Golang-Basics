package main

import "fmt"

type Number interface {
	~int64 | float64 //~ is Approximation
}

type Numbers[T Number] []T

type CustomInt int64

func (ci CustomInt) isPositive() bool {
	return ci > 0
}

func main() {
	//showSum()
	//showContains()
	//showAny()
	//unionInterfaceAndType()
	typeApproximation()
}

func showSum() {
	floats := []float64{1.1, 1.2, 1.3}
	ints := []int64{1, 2, 3}
	fmt.Println(sum(ints))
	fmt.Println(sum(floats))
}

func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func showContains() {
	type Person struct {
		name      string
		age       int64
		jobTittle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("ints:", contains(ints, 4))

	strings := []string{"Masha", "Misha", "Max", "Karakal"}
	fmt.Println("strings:", contains(strings, "Max"))
	fmt.Println("strings:", contains(strings, "Uzbek"))

	people := []Person{
		{
			"Max",
			21,
			"Backend",
		},
		{
			"Lev",
			35,
			"University",
		},
	}
	fmt.Println("people:", contains(people, Person{"Max", 31, "Front-end"}))
	fmt.Println("people:", contains(people, Person{"Lev", 35, "University"}))
}

func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}

	return false
}

func showAny() {
	show(1, 2, 3, 4, 5)
	show("test1", "test2", "test3")
	show([]int64{1, 2, 3}, []int64{10, 11, 12})
	show(map[string]int64{
		"first":  1,
		"second": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct {
		name string
	}{name: "Vasya"}))
}

func show[T any](entities ...T) {
	fmt.Println(entities)
}

func unionInterfaceAndType() {
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	floats := Numbers[float64]{1.0, 2.0, 5}

	fmt.Println("Sum of ints: ", sumUnionInterface(ints))
	fmt.Println("Sum of floats: ", sumUnionInterface(floats))

}

func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func typeApproximation() {
	customInts := []CustomInt{1, 2, 3}
	castedInts := make([]int64, len(customInts))

	for idx, value := range customInts {
		castedInts[idx] = int64(value)
	}

	fmt.Println("Sum of custom ints:", sumUnionInterface(customInts)) //Works only if ~ before int on CustomInts interface
	fmt.Println("Sum of casted ints:", sumUnionInterface(castedInts))
}
