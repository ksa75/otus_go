package main

import (
	"fmt"
)

func main() {
	s := "hello"

	defer fmt.Println(s)

	s = "perfect"
	defer func(str string) {
		fmt.Println(str, " ", s)
	}(s)
	//defer func() {
	//	fmt.Println(s)
	//}()

	defer func(str string) func(string) {
		defer fmt.Println("Сработал внешний дефер")
		return func(str2 string) {
			defer fmt.Println("Сработал внутренний дефер")
			fmt.Println(str2)
		}
	}(s)(s)

	defer func(str string) func(string) {
		return func(str2 string) {
			fmt.Println(s)
		}
	}("")(s)
	s = "world"
}
