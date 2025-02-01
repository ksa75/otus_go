package main

import (
	"fmt"
	"github.com/go-loremipsum/loremipsum"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
