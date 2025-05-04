package exbench

func Fast() int {
	acc := new(int)
	for range 10 {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}

	return *acc
}

func Slow() int {
	acc := new(int)
	for range 1000 {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}

	return *acc
}
