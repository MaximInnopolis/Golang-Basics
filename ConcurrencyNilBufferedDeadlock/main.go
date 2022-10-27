package main

import (
	"fmt"
)

func main() {

	//nilChannel()
	//unbufferedChannel()
	//bufferedChannel()
	forRange()
}

func nilChannel() {
	var nilChannel chan int

	fmt.Printf("Length: %d, capacity: %d\n", len(nilChannel), cap(nilChannel))

	//write to nil channel blocks forever
	//nilChannel <- 1

	//read from nil channel blocks forever
	//<-nilChannel

	//close(nilChannel)
}

func unbufferedChannel() {
	unbufferedChannel := make(chan int)

	fmt.Printf("Length: %d, capacity: %d\n", len(unbufferedChannel), cap(unbufferedChannel))

	//blocks until some1 read
	//unbufferedChannel <- 1

	//blocks until some1 write
	//<-unbufferedChannel

	//blocks on reading then write
	//go func(chanForWriting chan <- int) {
	//	time.Sleep(time.Second)
	//	unbufferedChannel <- 1
	//}(unbufferedChannel)
	//
	//value := <-unbufferedChannel
	//fmt.Println(value)

	////blocks on writing then read
	//go func(chanForReading chan int) {
	//	time.Sleep(time.Second)
	//	value := <-unbufferedChannel
	//	fmt.Println(value)
	//}(unbufferedChannel)
	//
	//unbufferedChannel <- 2

	//close channel
	//go func() {
	//	time.Sleep(time.Second)
	//	close(unbufferedChannel)
	//}()
	//
	//go func() {
	//	time.Sleep(time.Second)
	//	value := <-unbufferedChannel
	//	fmt.Println(value)
	//}()

	//writing to closed channel causes panic
	//unbufferedChannel <- 3

	//close of closed channel lead to panic
	close(unbufferedChannel)
	close(unbufferedChannel)
}

func bufferedChannel() {
	bufferedChannel := make(chan int, 2)

	fmt.Printf("Length: %d, capacity: %d\n", len(bufferedChannel), cap(bufferedChannel))

	//does not block while buffer is not full
	bufferedChannel <- 1
	bufferedChannel <- 2
	fmt.Printf("Length: %d, capacity: %d\n", len(bufferedChannel), cap(bufferedChannel))

	//blocks if buffer is full
	//bufferedChannel <- 3

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Printf("Length: %d, capacity: %d\n", len(bufferedChannel), cap(bufferedChannel))

	//blocks to read if buffer is empty(nothing to read)
	//<-bufferedChannel
}

func forRange() {
	channel := make(chan int, 3)
	numbers := []int{5, 6, 7, 8}

	//show all elements with for
	go func() {
		for _, num := range numbers {
			channel <- num
		}
		close(channel)
	}()

	for {
		value, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(value)
	}

	fmt.Println("~~~~~~~~~~~~")

	//show with for range buffered
	bufferedChannel := make(chan int, 3)
	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for value := range bufferedChannel {
		fmt.Println(value)
	}

}
