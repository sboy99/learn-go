package database

import (
	"fmt"

	"learn/go-sql/exceptions"
)

func Query() {
	rows, err := connection.Query("SELECT * FROM albums")
	exceptions.HandleQueryException(err)

	for rows.Next() {
		var id int
		var title, artist string
		var price float64

		err := rows.Scan(&id, &title, &artist, &price)
		exceptions.HandleQueryException(err)

		// Process each row
		fmt.Printf("ID: %d, Title: %s, Artist: %s, Price: %.2f\n", id, title, artist, price)

	}

	rows.Close()
}
