package authors

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// HandleAuthors handle requests to the /api/v1/authors route.
func HandleAuthors(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		getAllAuthors(w, r)
		return
	case r.Method == http.MethodPost:
		createAuthor(w, r)
		return
	}

	log.Println("error: author ", http.StatusMethodNotAllowed)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}

// HandleAuthor handle requests to the /api/v1/authors/ route.
func HandleAuthor(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	urlparam := strings.Split(path, "/api/v1/authors/")[1]

	if urlparam == "" {
		http.Redirect(w, r, "/api/v1/authors", http.StatusSeeOther)
		return
	}

	idparam, err := strconv.Atoi(urlparam)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	switch {
	case r.Method == http.MethodGet:
		getAuthorByID(w, r, idparam)
		return
	case r.Method == http.MethodPut:
		updateAuthor(w, r, idparam)
		return
	case r.Method == http.MethodDelete:
		deleteAuthor(w, r, idparam)
		return
	}

	log.Println("error: author/:id ", http.StatusMethodNotAllowed)
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	return
}
