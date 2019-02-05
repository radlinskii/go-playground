package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

const sessionLength int = 30

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

func init() {
	tpl = template.Must(template.ParseGlob("myhttp/sessions/templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if ifLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		sID := uuid.NewV4()
		c := &http.Cookie{Name: "session", Value: sID.String()}
		http.SetCookie(w, c)

		dbSessions[c.Value] = session{un, time.Now()}

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, r}
		dbUsers[un] = u

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if ifLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password are invalid", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password are invalid", http.StatusForbidden)
			return
		}

		sID := uuid.NewV4()
		c := &http.Cookie{Name: "session", Value: sID.String()}
		http.SetCookie(w, c)

		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !ifLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{Name: "session", Value: "", MaxAge: -1}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * time.Duration(sessionLength)) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func cleanSessions() {
	fmt.Println("Before clean")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * time.Duration(sessionLength)) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("After clean")
	showSessions()
}

func showSessions() {
	fmt.Println("*****")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("***")
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !ifLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{Name: "session", Value: sID.String()}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	var u user
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func ifLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value].un
	_, ok := dbUsers[s]
	return ok
}
