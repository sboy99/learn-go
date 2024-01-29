package models

type User struct {
	Model
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Email    string `gorm:"index:unique_email; uniqueIndex" json:"email"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}
