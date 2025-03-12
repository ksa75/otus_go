package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	onceInit int32 = iota
	onceDoing
	onceDone
)

type Once struct {
	state int32
}

func (o *Once) Do(fn func()) {
	if atomic.CompareAndSwapInt32(&o.state, onceInit, onceDoing) {
		fn()
		atomic.StoreInt32(&o.state, onceDone)
		return
	}

	for atomic.LoadInt32(&o.state) == onceDoing {
		time.Sleep(10 * time.Nanosecond)
	}
}

func main() {
	var once Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(onceBody)
			wg.Done()
		}()
	}
	wg.Wait()
}
