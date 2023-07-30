package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	conn := "host=localhost user=postgres password=root dbname=poketeams port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatalf("Error in database connection")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error in db ping")
	}

	return db
}
