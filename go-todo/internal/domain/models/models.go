package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

type User struct {
	Model
	Name   string  `json:"name"`
	Slug   string  `json:"slug"`
	Avatar *string `json:"avatar"`
	Email  string  `json:"email"`
	Pass   string  `json:"pass"`
	// Tasks  []Task  `json:"tasks"`
}

type Task struct {
	Model
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	UserId   uint   `json:"userId"`
}
