package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type middleware func(http.Handler) http.Handler

func useLogMiddleware() middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before1")
			defer log.Println("after1")
			h.ServeHTTP(w, r)
		})
	}
}

func useCustomLoggerMiddleware(logger *log.Logger) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("before2")
			defer logger.Println("after2")
			h.ServeHTTP(w, r)
		})
	}
}

func applyMiddlewares(h http.Handler, middlewares ...middleware) http.Handler {
	if len(middlewares) == 0 {
		panic("applyMiddlewares used withour a reason")
	}

	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func main() {
	logger := log.New(os.Stdout, "server: ", log.Lshortfile|log.Ldate|log.Lmicroseconds)

	http.Handle("/", applyMiddlewares(http.HandlerFunc(indexHandler), useLogMiddleware()))
	http.Handle("/foo", applyMiddlewares(http.HandlerFunc(fooHandler), useLogMiddleware(), useCustomLoggerMiddleware(logger)))
	http.HandleFunc("/foo/", bazzHandler) // /foo/:id
	http.HandleFunc("/bar", barHandler)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Write([]byte("hello index!"))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Write([]byte("hello foo!"))
}

func bazzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	path := r.URL.Path
	urlparam := strings.Split(path, "/foo/")[1]

	if urlparam == "" {
		http.Redirect(w, r, "/foo", http.StatusPermanentRedirect) // send "/foo/" request to "/foo" handler
		return
	}

	idparam, err := strconv.Atoi(urlparam)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Println(idparam)
	w.Write([]byte("well, hello there id = " + strconv.Itoa(idparam) + "!"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/foo/")
	if err != nil {
		fmt.Println(err)
	}

	j := make([]byte, resp.ContentLength)

	defer resp.Body.Close()

	_, err = resp.Body.Read(j)
	if err != io.EOF {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}
