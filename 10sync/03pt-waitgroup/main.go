package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const goCount = 5
	wg := sync.WaitGroup{}
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func(wg1 *sync.WaitGroup) {
			fmt.Println(i)
			time.Sleep(time.Second * 5)
			wg1.Done() // <===
		}(&wg)
	}
	wg.Wait() // <===
	fmt.Println("DONE")
}
