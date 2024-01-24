package exceptions

import (
	"log"
)

func HandleDatabaseConnectionException(err error) {
	if err != nil {
		log.Fatal("Database: ", err)
	}
}
