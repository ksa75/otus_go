package main

import "fmt"

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Integer interface {
	Signed | Unsigned
}

// Напишите констрейнт Number, который в себя будет включать все возможные числа
// Сделайте так, чтобы функция Double работала со всеми описанными в main случаями

type Price float32

func main() {
	fmt.Println(Double(0.2))
	fmt.Println(Double(uint16(2)))
	fmt.Println(Double(3))
	fmt.Println(Double(Price(4)))
}

func Double(x int) int {
	return x * x
}
