package main

// пример, когда дженерик-функция не принимает параметров дженерик-типа
// а возвращает значение типа, производного от T (слайс - производный тип)

func sliceFactory[T any](n int) []T {
	return make([]T, n)
}
