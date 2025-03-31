package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1) // <===
		go func() {

			defer wg.Done() // <===
			time.Sleep(time.Second * 5)
			fmt.Println(i)

		}()
	}
	wg.Wait()

}
