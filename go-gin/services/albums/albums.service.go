package albums

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAlbums() []Album {
	albums := []Album{
		{Id: 1, Title: "Sample Title", Artist: "Sample Artist", Price: 55},
		{Id: 2, Title: "Second Sample Title", Artist: "second Sample Artist", Price: 58},
		{Id: 3, Title: "Third Sample Title", Artist: "Third Sample Artist", Price: 95},
	}

	return albums
}
