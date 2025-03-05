package main

import "fmt"

func main() {

	sl := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("Длина %d , емкость %d \n", len(sl), cap(sl))
	sl = append(sl, 7)
	fmt.Printf("Длина %d , емкость %d \n", len(sl), cap(sl))

	sl2 := append(sl, 8)
	fmt.Printf("Длина sl %d , емкость %d \n", len(sl), cap(sl))
	fmt.Println("Слайс sl ", sl)
	fmt.Printf("Длина sl2 %d , емкость %d \n", len(sl2), cap(sl2))
	fmt.Println("Слайс sl2 ", sl2)

	sl[0] = 100
	fmt.Printf("sl2[0] %d ", sl2[0])
	fmt.Printf("Длина sl %d , емкость %d \n", len(sl), cap(sl))
	fmt.Println("Слайс sl ", sl)
	fmt.Printf("Длина sl2 %d , емкость %d \n", len(sl2), cap(sl2))
	fmt.Println("Слайс sl2 ", sl2)

	sl = append(sl, 10)
	fmt.Printf("sl2[0] %d \n", sl2[0])
	fmt.Printf("Длина sl %d , емкость %d, адрес %p \n", len(sl), cap(sl), &sl[0])
	fmt.Println("Слайс sl ", sl)
	fmt.Printf("Длина sl2 %d , емкость %d, адрес %p \n", len(sl2), cap(sl2), &sl2[0])
	fmt.Println("Слайс sl2 ", sl2)

	subslice := sl[5:7:7]
	subslice[0] = 0
	fmt.Printf("Длина sl %d , емкость %d \n", len(sl), cap(sl))
	fmt.Println("Слайс sl ", sl)
	fmt.Printf("Длина sl2 %d , емкость %d \n", len(sl2), cap(sl2))
	fmt.Println("Слайс sl2 ", sl2)

	fmt.Printf("Длина subslice %d , емкость %d \n", len(subslice), cap(subslice))
	fmt.Println("Слайс sl2 ", subslice)
	subslice = append(subslice, 200, 300)
	fmt.Printf("Длина sl2 %d , емкость %d \n", len(sl2), cap(sl2))
	fmt.Println("Слайс sl2 ", sl2)

	subslice2 := sl[3:5:6]
	fmt.Println("Сабслайс 2: ", subslice2)
	fmt.Println("Слайс sl2 ", sl2)
	appnd(subslice2)
	fmt.Println("Сабслайс 2: ", subslice2)
	fmt.Println("Слайс sl2 ", sl2)

}

func appnd(sl []int) {
	sl = append(sl, 2025)
}

func Experiment() {
	fmt.Println("Начинаем эксперимент со слайсами")

}
