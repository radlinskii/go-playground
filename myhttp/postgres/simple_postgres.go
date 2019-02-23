package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	http.HandleFunc("/api/v1/authors", getAuthors)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAuthors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("error: authors ", http.StatusMethodNotAllowed)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	q := `SELECT * from authors where age < $1 ;`

	rows, err := db.Query(q, 27)
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
