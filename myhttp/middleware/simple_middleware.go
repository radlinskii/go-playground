package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type middleware func(http.Handler) http.Handler

func useGZipMiddleware() middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before gzip")
			encodings := r.Header.Get("Accept-Encoding")
			if !strings.Contains(encodings, "gzip") {
				h.ServeHTTP(w, r)
				return
			}
			w.Header().Add("Content-Encoding", "gzip")
			gw := gzip.NewWriter(w)
			defer gw.Close()
			grw := gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gw,
			}
			h.ServeHTTP(grw, r)
			log.Println("after gzip")
		})
	}
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	log.Println("gzip")
	log.Println(len(data))
	if len(data) > 1400 {
		return grw.Writer.Write(data)
	}
	grw.Header().Del("Content-Encoding")
	grw.Header().Add("Content-Length", strconv.Itoa(len(data)))
	return grw.ResponseWriter.Write(data)
}

func useDurationMiddleware() middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before duration")
			t1 := time.Now()
			h.ServeHTTP(w, r)
			log.Println(time.Now().Sub(t1))
			log.Println("after duration")
		})
	}
}

func useCustomLoggerMiddleware(logger *log.Logger) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("before logger")
			logger.Printf("%s %s \n", r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
			log.Println("after logger")
		})
	}
}

func applyMiddlewares(h http.Handler, middlewares ...middleware) http.Handler {
	if len(middlewares) == 0 {
		panic("applyMiddlewares used withour a reason")
	}

	l := len(middlewares)
	for i := range middlewares {
		h = middlewares[l-i-1](h)
	}
	return h
}

func main() {
	logger := log.New(os.Stdout, "server: ", log.Lshortfile|log.Ldate|log.Lmicroseconds)

	http.Handle("/", applyMiddlewares(http.HandlerFunc(indexHandler), useDurationMiddleware()))
	http.Handle("/foo", applyMiddlewares(http.HandlerFunc(fooHandler), useDurationMiddleware(), useCustomLoggerMiddleware(logger)))
	http.HandleFunc("/foo/", bazzHandler) // /foo/:id
	http.HandleFunc("/bar", barHandler)
	http.Handle("/get/doggo-with-compression", applyMiddlewares(http.HandlerFunc(getDoggo), useDurationMiddleware(), useCustomLoggerMiddleware(logger), useGZipMiddleware()))
	http.Handle("/get/doggo-without-compression", applyMiddlewares(http.HandlerFunc(getDoggo), useDurationMiddleware(), useCustomLoggerMiddleware(logger)))

	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("hello index!"))
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello foo!"))
}

func bazzHandler(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte("fetched: " + string(j)))
}

func getDoggo(w http.ResponseWriter, r *http.Request) {

	type doggo struct {
		Name  string
		Breed string
		Age   int
	}

	w.Header().Set("Content-Type", "application/json")
	doggos := make([]doggo, 0)

	for i := 0; i < 33; i++ { // test gzip with 33 | 34 iterations
		doggos = append(doggos, doggo{"doggo" + strconv.Itoa(i), "mutt", i%6 + 1})
	}

	json, err := json.Marshal(doggos)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}
