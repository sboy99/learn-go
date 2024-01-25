package database

import (
	"encoding/json"

	"learn/go-sql/exceptions"
	"learn/go-sql/file"
)

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func Query() {
	var albums []Album
	rows, err := connection.Query("SELECT * FROM albums")
	exceptions.HandleQueryException(err)

	for rows.Next() {
		var album Album

		err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		exceptions.HandleQueryException(err)

		// Process each row
		albums = append(albums, album)
	}

	rows.Close()

	jsonAlbums, err := json.Marshal(albums)
	exceptions.HandleBasicException(err)

	file.Save("output.json", jsonAlbums)
}
