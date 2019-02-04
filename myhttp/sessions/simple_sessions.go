package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			t := time.Now()
			id := uuid.NewV4()
			cookie = &http.Cookie{Name: "session", Value: id.String(), Expires: t.Add(30 * 24 * time.Hour)}
			http.SetCookie(w, cookie)
		} else {
			http.Error(w, "Server Error!", http.StatusInternalServerError)
			_ = fmt.Errorf("Server Error: %s", err.Error())
			return
		}
	}
	fmt.Fprintln(w, cookie)
}
