package main

// мы можем внутри дженерик-функции создавать новые экземпляры T-типа,
// но только если во всех возможных инстансах это Value-type!

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

func create[T interface{ something | str }]() T {
	v := T{}

}
