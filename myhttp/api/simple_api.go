package main

import (
	"log"
	"net/http"

	"github.com/radlinskii/go-playground/myhttp/api/authors"
	"github.com/radlinskii/go-playground/myhttp/api/config"
	"github.com/radlinskii/go-playground/myhttp/api/middlewares"
)

// TODO separate middlewares for different method handlers on same route. Router??
func main() {
	http.Handle("/api/v1/authors", middlewares.Apply(http.HandlerFunc(authors.HandleAuthors),
		middlewares.UseDuration(config.Logger),
		middlewares.UseGZip(),
	))
	http.Handle("/api/v1/authors/", middlewares.Apply(http.HandlerFunc(authors.HandleAuthor),
		middlewares.UseDuration(config.Logger),
		middlewares.UseGZip(),
	))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", middlewares.UseCustomLogger(config.Logger)(http.DefaultServeMux)))
}
