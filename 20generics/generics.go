package main

import "fmt"

// комплексный пример, иллюстрирующий
// 1) соответствие структуры интерфейсу (классика)
// 2) соответствие дженерик-структуры интерфейсу (только при некоторых инстанциациях)
// 3) соответствие обычной структуры дженерик-интерфейсу (только при некоторых инстанциациях)
// 4) соответствие дженерик-структуры дженерик-интерфейсу (только при согласованных типах инстанцирования)
// 5) инстанцирование дженерик-типа другим дженерик-типом (дженерик структуры дженерик интерфейсом)

type Valuer interface {
	GetValue() int
}

type ValuerGene[T any] interface {
	GetValue() T
}

type Container[T any] struct {
	Value T
}

func (c Container[T]) GetValue() T {

	return c.Value
}

func main() {
	var v Valuer

	cont := Container[int]{Value: 10}
	v = cont
	cont2 := Container[string]{Value: "aaa"}

	fmt.Println(v.GetValue())

	var v2 ValuerGene[int]

	v2 = &str{}

	v2 = cont
	fmt.Println(v2.GetValue())

	var v3 ValuerGene[string]
	v3 = &something{}

	v3 = cont2
	_ = v3
	cont3 := Container[ValuerGene[string]]{}

	cont3.Value = &something{}
	cont3.Value = cont2

	cont4 := Container[string]{Value: "Result"}
	v3 = cont4
}

type something struct {
}

func (s *something) GetValue() string {
	return "Hello"
}

type str struct {
}

func (s *str) GetValue() int {
	return 100
}
