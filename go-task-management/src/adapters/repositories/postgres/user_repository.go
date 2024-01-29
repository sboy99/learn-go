package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/sboy99/learn-go/go-task-management/src/internal/domain/models"
)

// user repository structure
type UserPostgresRepository struct {
	DB *gorm.DB
}

func (r *UserPostgresRepository) Create(name string, email string) *models.User {
	newUser := models.User{
		Name:  name,
		Email: email,
	}
	r.DB.Create(&newUser)
	return &newUser
}
