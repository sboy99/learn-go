package ports

import (
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
)

// -------------------------------INTERFACES--------------------------------- //

type IUserRepositoryPort interface {
	Create(name string, slug string) (*models.User, error)
	GetSlug(name string, delimiter *string) (*string, error)
}

type IUserServicePort interface {
	Create(name string) (*models.User, error)
}
