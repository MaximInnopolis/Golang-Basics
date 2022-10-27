package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//
	//fmt.Println(sum(2, 3))
	//deferValues()

	//runtime.GOMAXPROCS(1) //number of parallel goroutines can be set
	//fmt.Println(runtime.NumCPU())
	go showNumbers(100)
	runtime.Gosched()       //switch to another goroutine
	time.Sleep(time.Second) //set number of seconds to sleep specific goroutine
	makePanic()
	fmt.Println("exit")
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func sum(x int, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y

	return
}

func deferValues() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("First", i)
	}

	for i := 0; i < 10; i++ { //Wrong usage of defer
		defer func() {
			fmt.Println("Second", i)
		}()
	}

	for i := 0; i < 10; i++ { //Correct version of Second
		k := i
		defer func() {
			fmt.Println("Third", k)
		}()
	}

	for i := 0; i < 10; i++ { //Correct version of Second
		defer func(k int) {
			fmt.Println("Fourth", k)
		}(i)
	}
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("Some panic")
	fmt.Println("Unreachable code")
}
