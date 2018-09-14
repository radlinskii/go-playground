package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var templates = template.Must(
	template.ParseFiles(
		filepath.Join("tmpl", "edit.html"),
		filepath.Join("tmpl", "view.html")))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var validDataPath = regexp.MustCompile("^data/([a-zA-Z0-9]+).txt$")

type Page struct {
	Title    string
	Body     []byte
	HTMLBody template.HTML
}

func (p *Page) save() error {
	filename := filepath.Join("data", p.Title+".txt")
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) parseLinks() {
	titles := getPagesTitles()
	for i := range titles {
		p.Body = []byte(strings.Replace(string(p.Body), titles[i], "<a href=\"/view/"+titles[i]+"\">"+titles[i]+"</a>", -1))
	}
	p.HTMLBody = template.HTML(p.Body)
}

func getPagesTitles() []string {
	var files []string

	root := "data"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		m := validDataPath.FindStringSubmatch(path)
		if len(m) > 1 {
			files = append(files, m[1])
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func loadPage(title string) (*Page, error) {
	filename := filepath.Join("data", title+".txt")
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmplt string, p *Page) {
	err := templates.ExecuteTemplate(w, tmplt+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	p.parseLinks()
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, _ *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
