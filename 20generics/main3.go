package main

import "fmt"

// Задание констрейнта с ~
// вызов дженерик-функции с типами-синонимами и через приведение к типам-синонимам

type MyConstraint interface {
	~int | int32
}

type myOwnInt int
type myOwnInt2 myOwnInt

func main() {
	var a int = 10
	var b int = 20
	var c myOwnInt2 = 30
	var d myOwnInt2 = 30
	var res myOwnInt2

	_ = Max[int](a, b)
	res = myOwnInt2(Max[int](int(c), int(d)))
	res = Max[myOwnInt2](c, d)
	fmt.Println(res)
}

func Max[T MyConstraint](a, b T) T {
	if a > b {
		return a
	}
	return b
}
