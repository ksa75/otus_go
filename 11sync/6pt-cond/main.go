package main

import (
	"fmt"
	"sync"
	"time"
)

// Доработайте тип Worker
// с помощью sync.Cond.

type Worker struct {
	mu    *sync.Mutex
	cond  *sync.Cond
	ready bool
}

func NewWorker() *Worker {
	mx := &sync.Mutex{}
	w := &Worker{
		mu: mx,
		cond: &sync.Cond{
			L: mx,
		},
		ready: false,
	}
	return w
}

func (m *Worker) setReady() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ready = true

	m.cond.Broadcast()
}

func (m *Worker) checkReady() bool {
	return m.ready
}

func (m *Worker) WaitReady() {
	m.mu.Lock()
	for !m.checkReady() {
		m.cond.Wait()
	}
	// может быть что то еще полезное сделаем
	m.mu.Unlock()

}

func main() {
	w := NewWorker()
	go func() { // имитация инициализации
		time.Sleep(time.Second)
		fmt.Println("I'm ready")
		w.setReady()
	}()

	w.WaitReady()

	fmt.Println("Ready to use Worker")
}
