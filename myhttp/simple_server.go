package main

import (
	"fmt"
	"net/http"
)

func r1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Sunset! %d", 1)
}

func r2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Sunshine! %d", 2)
}

func main() {
	http.HandleFunc("/r1", r1)
	http.HandleFunc("/r2/", r2)

	http.ListenAndServe(":3000", nil)
}

