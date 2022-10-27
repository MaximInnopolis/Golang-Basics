package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	//chanAsPromise()
	//chanAsMutex()
	//withoutErrorGroup()
	errorGroup()
}

func makeRequest(num int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second / 2)
		responseChan <- fmt.Sprintf("Response number %d", num)
	}()

	return responseChan
}

func chanAsPromise() {
	firstResponseChan := makeRequest(1)
	secondResponseChan := makeRequest(2)
	// do something else

	fmt.Println("Non blocking")
	fmt.Println(<-firstResponseChan, <-secondResponseChan)
}

func chanAsMutex() {
	var counter int
	wg := sync.WaitGroup{}
	mutexChan := make(chan struct{}, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutexChan <- struct{}{}

			counter++

			<-mutexChan
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func withoutErrorGroup() {
	var err error
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("First started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Second started")
			err = fmt.Errorf("Any error")
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Third started")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err)
}

func errorGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
		default:
			fmt.Println("First started")
			time.Sleep(time.Second)
		}
		return nil
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("Second started")
			return fmt.Errorf("Any error")
		}
		return nil
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():

		default:
			fmt.Println("Third started")
			time.Sleep(time.Second)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
