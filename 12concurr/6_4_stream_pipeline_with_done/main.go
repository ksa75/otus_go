package main

import "fmt"

func main() {
	generator := func(done <-chan struct{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan struct{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan struct{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan struct{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	chOutMul1 := multiply(done, intStream, 2)
	chOutAdd1 := add(done, chOutMul1, 1)
	resChan := multiply(done, chOutAdd1, 10)

	//pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range resChan {
		fmt.Println(v)
	}
}
