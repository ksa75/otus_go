package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	s := []User{
		{"vasya", 19},
		{"petya", 18},
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Name > s[j].Name
	})

	fmt.Println(s)
}
