package main

import (
	"fmt"
	"sync"
)

type Counters struct {
	m sync.Map
}

func NewCounters() *Counters {
	return &Counters{}
}

func (c *Counters) Load(key string) (int, bool) {
	val, ok := c.m.Load(key)
	if !ok {
		return 0, false
	}
	return val.(int), true
}

func (c *Counters) Store(key string, value int) {
	c.m.Store(key, value)
}

func (c *Counters) Range(f func(k string, v int) bool) {
	c.m.Range(func(k, v interface{}) bool {
		return f(k.(string), v.(int))
	})
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
