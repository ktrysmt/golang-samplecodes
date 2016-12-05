package main

import (
	"fmt"
	"net/http"
)

var handle = http.HandlerFunc(handler)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/a", handle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
