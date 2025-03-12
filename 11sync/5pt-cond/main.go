package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var cond *sync.Cond
var tasks []func()

func worker(i int) {
	var task func()

	mu.Lock()
	//fmt.Printf("%d Прошел через lock \n", i)

	for len(tasks) == 0 { // ждем пока там
		cond.Wait()
		//fmt.Printf("%d Побежал дальше \n", i)
	}
	task, tasks = tasks[0], tasks[1:]
	mu.Unlock()

	task()
}

func produce(task func()) {
	mu.Lock()
	tasks = append(tasks, task)
	mu.Unlock()

	cond.Signal()
}

func main() {
	cond = sync.NewCond(&mu)

	for i := 0; i < 5; i++ {

		go worker(i)
	}

	produce(func() { fmt.Println("1") })

	time.Sleep(time.Second)
	produce(func() { fmt.Println("2") })
	time.Sleep(time.Second)
	produce(func() { fmt.Println("3") })
	time.Sleep(time.Second)

	time.Sleep(time.Second)
}
