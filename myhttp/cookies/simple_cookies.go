package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/increment", increment)
	http.HandleFunc("/decrement", decrement)
	http.HandleFunc("/delete", delete)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	fmt.Fprintln(w, `<h1>Home</h1>`)
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, `<h2>No counter cookie found!</h2>`)
		} else {
			http.Error(w, "Server Error!", http.StatusInternalServerError)
			_ = fmt.Errorf("Server Error: %s", err.Error())
			return
		}
	} else {
		count, err := strconv.Atoi(cookie.Value)
		if err != nil {
			http.Error(w, "Server Error!", http.StatusInternalServerError)
			_ = fmt.Errorf("Server Error: %s", err.Error())
			return
		}

		fmt.Fprintf(w, `<h2>counter: %d</h2>`, count)
	}

	fmt.Fprintln(w, `<a href="/increment"> increment cookie </a><br/>
	<a href="/decrement"> decrement cookie </a><br/>
	<a href="/"> check the home page </a><br/>
	<a href="/delete"> delete a cookie </a><br/>`)
}

func increment(w http.ResponseWriter, req *http.Request) {
	addValueToCounterCookie(&w, req, 1)
}

func decrement(w http.ResponseWriter, req *http.Request) {
	addValueToCounterCookie(&w, req, -1)
}

func addValueToCounterCookie(w *http.ResponseWriter, req *http.Request, value int) {
	cookie, err := req.Cookie("counter")
	if err != nil {
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{Name: "counter", Value: "0", MaxAge: 3600}
		} else {
			http.Error(*w, "Server Error!", http.StatusInternalServerError)
			_ = fmt.Errorf("Server Error: %s", err.Error())
			return
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(*w, "Server Error!", http.StatusInternalServerError)
		_ = fmt.Errorf("Server Error: %s", err.Error())
		return
	}

	count += value

	cookie.Value = strconv.Itoa(count)
	http.SetCookie(*w, cookie)

	fmt.Fprintf(*w, `<h1>incremental page</h1>
	<h2>counter: %s</h2>
	<a href="/increment"> increment cookie </a><br/>
	<a href="/decrement"> decrement cookie </a><br/>
	<a href="/"> check the home page </a><br/>
	<a href="/delete"> delete a cookie </a><br/>`, cookie.Value)
}

func delete(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	cookie.MaxAge = -1

	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
