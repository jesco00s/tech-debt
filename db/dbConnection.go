package db

import (
	"database/sql"

	"log"
)

func GetDbConnection() *sql.DB {
	connStr := "user=postgres dbname=josejescobar host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
