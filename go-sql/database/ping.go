package database

import (
	"fmt"

	"learn/go-sql/exceptions"
)

func Ping() {
	err := connection.Ping()
	exceptions.HandleDatabaseConnectionException(err)
	fmt.Println("Conneted to database")
}
