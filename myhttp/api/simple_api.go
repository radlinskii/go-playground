package main

import (
	"log"
	"net/http"

	"github.com/radlinskii/go-playground/myhttp/api/authors"
)

func main() {
	http.HandleFunc("/api/v1/authors/", authors.HandleAuthors)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
