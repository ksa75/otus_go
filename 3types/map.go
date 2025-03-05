package main

import "fmt"

func main() {

	m := map[string]int{}
	//count := m["Петров"]
	//count++
	//m["Петров"] = count

	m["Петров"]++
	fmt.Println(m["Петров"])

	m2 := make(map[string]struct{})

	m2["Петров"] = struct{}{}
	if _, ok := m2["Петров"]; ok {
		fmt.Println("Да, есть у нас Петров")
	} else {
		fmt.Println("Нет, Петрова у нас нет")
	}

}
