package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//baseKnowledge()
	workerPool()
}

func baseKnowledge() {
	ctx := context.Background()
	fmt.Println(ctx)

	toDO := context.TODO()
	fmt.Println(toDO)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~")

	withValue := context.WithValue(ctx, "name", "Max")
	fmt.Println(withValue.Value("name"))

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~")

	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~")

	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*2))
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~")

	withTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	fmt.Println(<-withTimeout.Done())
}

func workerPool() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	wg := sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			//if i == 500 {
			//	cancel()
			//}
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for number := range processedNumbers {
		counter++
		fmt.Println(number)
	}
	fmt.Println(counter)
}

func work(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}
