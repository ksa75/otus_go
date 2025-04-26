package main

// Создайте универсальную функцию, которая принимает срез любого типа,
// функцию преобразования элементов и возвращает новый срез с преобразованными элементами,
// используя дженерики.
func TransformInt(sl []int, f func(i int) int) []int {
	//res:=make([]int, 0, len(sl))
	//for _, v:=range sl {
	//  res=append(res,  f(v))
	//}
	res := make([]int, len(sl))
	for i, v := range sl {
		res[i] = f(v)
	}
	return res

}

func TransformString(sl []string, f func(s string) string) []string {
	res := make([]string, len(sl))
	for i, v := range sl {
		res[i] = f(v)
	}
	return res
}

func TransformAny[T, V any](sl []T, f func(t T) V) []V {
	res := make([]V, len(sl))
	for i, v := range sl {
		res[i] = f(v)
	}
	return res
}
