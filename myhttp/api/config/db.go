package config

import (
	"database/sql"
	"fmt"
	"log"

	// postgres driver
	_ "github.com/lib/pq"
)

// DB is a pointer to sql.DB object.
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://ignacyradlinski:@localhost/iconomy?sslmode=disable") // TODO .env file?
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database.")
}
