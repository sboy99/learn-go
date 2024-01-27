package models

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
}
