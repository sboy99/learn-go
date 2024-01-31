package models

import (
	"time"
)

type model struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt"`
}

type User struct {
	model
	Slug     string  `json:"slug"`
	Name     string  `json:"name"`
	Avatar   *string `json:"avatar"`
	IsBanned bool    `json:"isBanned"`
	Roles    []Role  `json:"roles"`
}

type Role struct {
	model
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
