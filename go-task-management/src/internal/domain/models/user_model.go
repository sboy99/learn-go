package models

type User struct {
	Model
	Name  string `json:"name"`
	Email string `gorm:"index:unique_email" json:"email"`
	Tasks []Task `json:"tasks"`
}
