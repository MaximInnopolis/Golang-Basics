package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//baseSelect()
	gracefulShutdown()
}

func baseSelect() {
	bufferedChan := make(chan string, 1)
	bufferedChan <- "first"

	select {
	case str := <-bufferedChan:
		fmt.Println("read", str)
	case bufferedChan <- "second":
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

	unbufferedChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbufferedChan <- 1
	}()

	select {
	case bufferedChan <- "third":
		fmt.Println("unblocked writing")
	case value := <-unbufferedChan:
		fmt.Println("blocked reading", value)
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("time is up")
	default:
		fmt.Println("default case")
	}

	resultChan := make(chan int)
	timer := time.After(time.Second) // timer outside loop

	go func() {
		defer close(resultChan)
		for i := 1; i < 1000; i++ {
			select {
			case <-timer:
				fmt.Println("time is up")
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for value := range resultChan {
		fmt.Println(value)
	}

}

func gracefulShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("time is up")
		return
	case sign := <-signalChan:
		fmt.Println("Stopped by signalЖ", sign)
		return
	}
}
