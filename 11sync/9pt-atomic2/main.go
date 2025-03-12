package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	var luckyNumber int32 = 7
	wg.Add(2)
	go func() {
		atomic.StoreInt32(&luckyNumber, 42)
		wg.Done()
	}()
	go func() {
		fmt.Println(atomic.LoadInt32(&luckyNumber))
		wg.Done()
	}()
	wg.Wait()
}
