package main

import (
	"fmt"
)

func main() {
	data, err := callServer()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Нет никакой ошибки")
		}
	}()

	if err = saveToDB(data); err != nil {
		//fmt.Println(err)
		return
	}

	return
}

func callServer() (int, error) {
	return 0, nil
}

func saveToDB(a int) error {
	return fmt.Errorf("save error")
}
