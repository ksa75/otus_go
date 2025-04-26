package main

import "fmt"

// Способы вызова дженерик-функции с большим количеством задаваемых типов
// можем инстанцировать дженерик-функцию и сохранить её в переменную функционального типа
// или можем инстанцировать прямо в месте вызова

type Numbers interface {
	int | int32 | int64 | float64
}

func main() {
	oldMap := map[string]float64{
		"one":  1.2,
		"two":  2.4,
		"four": 4.8,
	}
	//convert := ConvertMap[string, float64, int64]
	//newMap := convert(oldMap)

	//newMap:=ConvertMap[string, float64, int64](oldMap)

	var newMap map[string]int64
	newMap = ConvertMap[string, float64, int64](oldMap)

	var convert func(map[string]float64) map[string]int64
	convert = ConvertMap[string, float64, int64]
	newMap2 := convert(oldMap)

	fmt.Println("%T = %+v", newMap, newMap)
	fmt.Println("%T = %+v", newMap2, newMap2)
}

func ConvertMap[K comparable, FROM, TO Numbers](in map[K]FROM) map[K]TO {
	nMap := make(map[K]TO, len(in))
	for k, v := range in {
		nMap[k] = TO(v)
	}
	return nMap
}
