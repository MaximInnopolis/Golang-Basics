package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//WaitGroup

	//withoutWait()
	//withWait()
	//wrongAdd()

	//Mutex
	//writeWithoutConcurrency()
	//writeWithoutMutex()
	//writeWithMutex()
	//readWithMutex()
	readWithRWMutex()
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //Add n tasks to the wg
		go func(i int) {
			fmt.Println(i + 1)
			wg.Done() //Remove 1 task from the wg
		}(i)
	}

	wg.Wait() //As soon as in wg all tasks completed
	//it gives access to the following commands
	fmt.Println("exit")
}

func wrongAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(i int) { //There could be a situation in which
			//goroutine does not have enough time to
			//execute(add task in wg correspondingly)
			wg.Add(1)
			defer wg.Done()
			fmt.Println(i + 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

func writeWithoutConcurrency() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++ //At the moment of the time
			// it has to be for only 1 goroutine(read or write)
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			_ = counter
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var rwmu sync.RWMutex

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			rwmu.RLock()
			_ = counter
			rwmu.RUnlock()
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			rwmu.Lock()
			counter++
			rwmu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
