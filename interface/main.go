package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	name string
}

type Duck struct {
	fam string
}

func (s Human) Run() string {
	return "Human " + s.name + " is running"
}

func (s Human) writeCode() {
	fmt.Println("This is human " + s.name + "!!!")
}

func (s Duck) Run() string {
	return "Duck " + s.fam + " is running"
}

func (s Duck) writeCode() {
	fmt.Println("This is duck " + s.fam + "!!!")
}

func (s Duck) Fly() string {
	return "Duck " + s.fam + " fly"
}

func main() {
	var runner Runner
	fmt.Printf("Type = %T, value = %v\n", runner, runner)

	if runner == nil {
		fmt.Println("Runner is nil")
	}

	var unnamedperson *Human

	runner = unnamedperson
	fmt.Printf("Type = %T, value = %v\n", runner, runner)
	if runner != nil {
		fmt.Println("Runner isn't nil")
	}

	namedperson := &Human{"Max"}
	runner = namedperson
	fmt.Printf("Type = %T, value = %v\n", runner, runner)
	if runner != nil {
		fmt.Println("Runner isn't nil")
	}

	var emptyInterface interface{} = unnamedperson
	fmt.Printf("Type = %T, value = %v\n", emptyInterface, emptyInterface)

	fmt.Println("Polymorphism:")
	duck := &Duck{"Donald"}
	polymorphism(duck)
	polymorphism(namedperson)

	fmt.Println("\nType assertion")
	typeAssertion(namedperson)
	typeAssertion(duck)
}

func polymorphism(runner Runner) {
	fmt.Println(runner.Run())
}

func typeAssertion(runner Runner) {
	if human, ok := runner.(*Human); ok {
		human.writeCode()
	}

	if duck, ok := runner.(*Duck); ok {
		duck.writeCode()
	}

	switch v := runner.(type) {
	case *Human:
		fmt.Println(v.Run())
	case *Duck:
		fmt.Println(v.Fly())
	default:
		fmt.Printf("Type = %T, value = %v\n", v, v)
	}
}
