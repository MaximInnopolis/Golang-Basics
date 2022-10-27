package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	//addAtomic()
	//storeLoadSwap()
	//compareAndSwap()
	//atomicVal()
}

func addAtomic() { //faster than mutex
	start := time.Now()
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Counter = ", counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func storeLoadSwap() {
	var counter int64

	fmt.Println("Initial value :", atomic.LoadInt64(&counter))

	atomic.StoreInt64(&counter, 5)
	fmt.Println("Value after store:", atomic.LoadInt64(&counter))

	fmt.Println("Value after loaded:", atomic.SwapInt64(&counter, 10))
	fmt.Println("Value after swap stored:", atomic.LoadInt64(&counter))
}

func compareAndSwap() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}
			fmt.Println("Swapped goroutine number is :", i)
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}

func atomicVal() {
	var value atomic.Value
	value.Store(1)
	fmt.Println(value.Load())

	fmt.Println(value.Swap(2))
	fmt.Println(value.Load())

	fmt.Println(value.CompareAndSwap(2, 3))
}
