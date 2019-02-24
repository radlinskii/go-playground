package authors

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/radlinskii/go-playground/myhttp/api/config"
)

type author struct {
	ID          int    `json:"id"`
	Fname       string `json:"firstName"`
	Lname       string `json:"lastName"`
	Age         int    `json:"age"`
	Phone       string `json:"phoneNumber"`
	Description string `json:"description"`
}

// curl -i -X PUT -d "id=4&firstName=marcelina&lastName=radlinska&phoneNumber=+4865423322&age=25&description=author number 4" localhost:8080/api/v1/authors/5
func updateAuthor(w http.ResponseWriter, r *http.Request, id int) {
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

	result, err := config.DB.Exec(`UPDATE authors SET
	first_name = $1, last_name = $2, phone_number = $3, age = $4, description = $5
	WHERE author_id = $6;`, a.Fname, a.Lname, a.Phone, a.Age, a.Description, id)

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
		log.Println("No rows updated.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// curl -i -X DELETE localhost:8080/api/v1/authors/5
func deleteAuthor(w http.ResponseWriter, r *http.Request, id int) {
	result, err := config.DB.Exec(`DELETE FROM authors WHERE author_id = $1;`, id)

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
		log.Println("No rows deleted.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
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

	result, err := config.DB.Exec(`INSERT INTO authors
	(first_name, last_name, phone_number, age, description) VALUES
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

// curl -i -X GET localhost:8080/api/v1/authors/
func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := `SELECT * FROM authors;`

	rows, err := config.DB.Query(q)
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

// curl -i -X GET localhost:8080/api/v1/authors/1
func getAuthorByID(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Set("Content-Type", "application/json")
	q := `SELECT * FROM authors WHERE author_id = $1 ;`

	row := config.DB.QueryRow(q, id)

	a := author{}
	err := row.Scan(&a.ID, &a.Fname, &a.Lname, &a.Phone, &a.Age, &a.Description)
	switch {
	case err == sql.ErrNoRows:
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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
