package main

import (
	"fmt"
	"sync"
	"time"
)

type task struct {
	name  string
	sleep time.Duration
}

func doJob(t task, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("task %q begin\n", t.name)
	time.Sleep(t.sleep)
	fmt.Printf("task %q end\n", t.name)
}

func main() {
	tasks := []task{
		{"fast", time.Second / 10},
		{"slow", time.Second},
		{"moderate", time.Second / 2},
	}

	wg := sync.WaitGroup{}
	for _, t := range tasks {
		wg.Add(1)
		go doJob(t, &wg)
	}

	wg.Wait()
}
