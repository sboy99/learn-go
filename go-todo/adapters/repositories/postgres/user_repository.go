package postgres

import (
	"errors"

	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
	"github.com/sboy99/learn-go/go-todo/internal/helpers"
	"gorm.io/gorm"
)

// ------------------------------STRUCTS---------------------------------- //

type UserPostgresRepository struct {
	DB *gorm.DB
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (r *UserPostgresRepository) Create(name string, slug string) (*models.User, error) {
	// create a new task
	user := &models.User{
		Name: name,
		Slug: slug,
	}
	// save the task in database
	err := r.DB.Create(user).Error
	return user, err
}

func (r *UserPostgresRepository) GetSlug(name string, delimiter *string) (*string, error) {
	user := new(models.User)
	slug := helpers.GenSlugStr(name, delimiter)
	d := ""

	if delimiter != nil {
		d = *delimiter
	}

	for true {
		err := r.DB.First(user, &models.User{
			Slug: slug,
		}).Error
		if err != nil {
			return &slug, nil
		}

		slug += d + helpers.GetSlugSalt(1000)

	}

	return nil, errors.New("Can't generate slug")
}
