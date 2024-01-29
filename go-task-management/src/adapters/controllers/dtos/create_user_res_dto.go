package dtos

import "time"

type CreateUserResDto struct {
	Id        uint      `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
