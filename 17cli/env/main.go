package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {

	env := os.Environ()                           // слайс строк "key=value"
	fmt.Println(slices.Contains(env, "USER=rob")) // USER=rob

	user, exists := os.LookupEnv("USER")
	if exists {
		fmt.Println(user) // rob
	}

	os.Setenv("CITY", "Ottawa") // установить
	// syscall.Unsetenv("CITY")
	city, cityExists := os.LookupEnv("CITY")
	fmt.Println(city, cityExists)

	//os.Unsetenv("CITY")                                 // удалить
	fmt.Println(os.ExpandEnv("$USER lives in ${CITY}")) // "шаблонизация"
}

// CITY=Ekaterinburg
// export CITY
// unset CITY
