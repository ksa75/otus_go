package main

// Простейший пример объявления и вызова Дженерик-функции
import "fmt"

type Numbers interface {
	int | int32 | int64 | float64
}

func Nmax[T Numbers](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a = 10
	var b = 20
	var c = 2.5
	var d = 4.2

	fmt.Println(Nmax[int](a, b))
	fmt.Println(Nmax[float64](c, d))

	fmt.Println(Nmax(a, b))
	fmt.Println(Nmax(c, d))
}
