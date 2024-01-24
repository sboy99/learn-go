package main

import "learn/go-sql/database"

func main() {
	// connect to database
	database.Connect()
	database.Query()
	database.Close()
}
