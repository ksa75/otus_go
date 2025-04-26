package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//pprof.StartCPUProfile()
		fmt.Fprintf(w, "Hello, world!\n\n")
		//pprof.StopCPUProfile()
	})

	http.ListenAndServe(":8080", nil)
}
