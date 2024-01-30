package models

type Session struct {
	Id           string `json:"id"`
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"user"`
}
