package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}

	var v int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&v, 1) // атомарный v++
			//v++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}
