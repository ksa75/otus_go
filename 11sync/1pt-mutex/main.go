package main

import (
	"fmt"
	"sync"
)

type Counters struct {
	mu sync.Mutex
	m  map[string]int
}

func NewCounters() *Counters {
	return &Counters{
		m: make(map[string]int),
	}
}

func (c *Counters) Load(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.m[key]
	return val, ok
}

func (c *Counters) Store(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *Counters) Range(f func(k string, v int) bool) {
	for k, v := range c.m {
		if !f(k, v) {
			return
		}
	}
}

func main() {
	counters := NewCounters()
	_, ok := counters.Load("otus")
	fmt.Println("has otus:", ok)

	counters.Store("otus", 27)
	counters.Store("habr", 42)
	counters.Range(func(k string, v int) bool {
		fmt.Println("key:", k, "val:", v)
		return true
	})
}
