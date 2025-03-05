package main

import "fmt"

func main() {

	//var m map[string]int
	m := map[string]int{} // то же, что m:=make(map[string]int)

	sl := []string{"Вася", "Петя", "Вася", "Коля", "Саша", "Федя"}

	for _, v := range sl {
		m[v]++
	}

	//fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, ": ", v)
	}
	count := m["Вася"]
	count++
	m["Вася"] = count

	m["Вася"]++
}
