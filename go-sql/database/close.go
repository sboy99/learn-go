package database

import (
	"fmt"

	"learn/go-sql/exceptions"
)

func Close() {
	err := connection.Close()
	exceptions.HandleDatabaseConnectionException(err)
	fmt.Println("Disconneted from database")
}
