package main

// Изменение порядка настраиваемых типов в функции
// теперь при вызове можно указать явно лишь один первый тип
// остальные два компилятор сам выведет на основе типа аргумента функции
func main() {
	oldMap := map[string]float64{
		"one":  1.2,
		"two":  2.4,
		"four": 4.8,
	}
	var newMap map[string]int64
	newMap = ConvertMap[](oldMap)
	_ = newMap

}

func ConvertMap[TO Numbers, K comparable, FROM Numbers](in map[K]FROM) map[K]TO {
	nMap := make(map[K]TO, len(in))
	for k, v := range in {
		nMap[k] = TO(v)
	}
	return nMap
}
