package main

import (
	"testing"
)

const gcFreq = 100_000

func BenchmarkWithPoolGC(b *testing.B) {
	pool := NewPersonPool()
	for i := 0; i < b.N; i++ {
		person := pool.Get()
		person.name = "Ivan"
		pool.Put(person)
		//if (i % gcFreq) == 0 {
		//	runtime.GC()
		//}
	}
}
func BenchmarkWithoutPoolGC(b *testing.B) {
	var gPerson *Person
	for i := 0; i < b.N; i++ {
		person := &Person{name: "Ivan"}
		gPerson = person
		//if (i % gcFreq) == 0 {
		//	runtime.GC()
		//}
	}
	_ = gPerson
}
