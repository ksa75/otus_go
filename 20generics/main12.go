package main

import "fmt"

// Написать функцию, принимающую два слайса, и возвращающую мапу, составленную
// из ключей первого слайса, и значений второго

func TwoSlicesToMap[K comparable, V any](s1 []K, s2 []V) map[K]V {

	resLen := Nmin(len(s1), len(s2))
	//resLen:=min(len(s1), len(s2))
	m := make(map[K]V, resLen)
	for i := 0; i < resLen; i++ {
		m[s1[i]] = s2[i]
	}
	return m
}

func Nmin[T Numbers](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Point int32
type Points []Point

func main() {
	s1 := Points{1, 2, 3, 4}
	s2 := []string{"one, two", "three"}
	m := TwoSlicesToMap(s1, s2)
	fmt.Println(m)
}
