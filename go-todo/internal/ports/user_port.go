package ports

import (
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
)

// -------------------------------INTERFACES--------------------------------- //

type IUserRepositoryPort interface {
	Create(email string, pass string, name string, slug string) (*models.User, error)
	GetSlug(name string) (string, error)
	GetUserWithEmail(email string) (*models.User, error)
	GetUser(conds *models.User) (*models.User, error)
}

type IUserServicePort interface{}
