package main

import (
	"fmt"
	"time"
)

// Сделайте тип Worker потокобезопасным
// с помощью mutex или rwmutex.
// Объясните свой выбор.

type Worker struct {
	ready bool
}

func (w *Worker) SetReady() {

	w.ready = true

}

func (w *Worker) CheckReady() bool {

	return w.ready

}

func main() {
	w := Worker{}

	go func() { // имитация инициализации
		time.Sleep(time.Second)
		fmt.Println("I'm ready")
		w.SetReady()
	}()

	for !w.CheckReady() {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Ready to use Worker")
}
