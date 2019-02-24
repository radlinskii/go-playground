package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://ignacyradlinski:@localhost/iconomy?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database.")
}

type author struct {
	ID          int    `json:"id"`
	Fname       string `json:"firstName"`
	Lname       string `json:"lastName"`
	Age         int    `json:"age"`
	Phone       string `json:"phoneNumber"`
	Description string `json:"description"`
}

func main() {
	defer db.Close()
	http.HandleFunc("/api/v1/authors", authors)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func authors(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		val := r.FormValue("id")
		if val == "" {
			getAllAuthors(w, r)
			return
		}

		id, err := strconv.Atoi(val)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		getAuthorByID(w, r, id)
		return
	case r.Method == http.MethodPost:
		createAuthor(w, r)
		return
	}

	log.Println("error: authors ", http.StatusMethodNotAllowed)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

//  curl -i -X POST -d "id=4&firstName=marcelina&lastName=radlinska&phoneNumber=+4865423322&age=25&description=author number 4" localhost:8080/api/v1/authors
func createAuthor(w http.ResponseWriter, r *http.Request) {
	a := author{}
	ageval := r.FormValue("age")
	a.Fname = r.FormValue("firstName")
	a.Lname = r.FormValue("lastName")
	a.Phone = r.FormValue("phoneNumber")
	a.Description = r.FormValue("description")

	if a.Fname == "" || a.Lname == "" || a.Phone == "" || a.Description == "" || ageval == "" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	age, err := strconv.Atoi(ageval)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	a.Age = age

	result, err := db.Exec(`INSERT INTO authors
	(first_name, last_name, phone_number, age, description) values
	($1, $2, $3, $4, $5);`, a.Fname, a.Lname, a.Phone, a.Age, a.Description)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows != 1 {
		log.Println("No rows inserted.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := `SELECT * from authors;`

	rows, err := db.Query(q)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	as := make([]author, 0)
	for rows.Next() {
		a := author{}
		if err := rows.Scan(&a.ID, &a.Fname, &a.Lname, &a.Phone, &a.Age, &a.Description); err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		as = append(as, a)
	}
	if err = rows.Err(); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(as)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func getAuthorByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	q := `SELECT * from authors where author_id = $1 ;`

	row := db.QueryRow(q, id)

	a := author{}
	err := row.Scan(&a.ID, &a.Fname, &a.Lname, &a.Phone, &a.Age, &a.Description)
	switch {
	case err == sql.ErrNoRows:
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	case err != nil:
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(a)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
