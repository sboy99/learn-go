package database

import (
	"database/sql"
	"os"

	"learn/go-sql/exceptions"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func Connect() {
	// "postgres://pqgotest:postgres@localhost/godb?sslmode=disable"
	connString := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", connString)
	exceptions.HandleDatabaseConnectionException(err)
	connection = db

	// check connection
	Ping()
}
