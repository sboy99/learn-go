package services

import (
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
	"github.com/sboy99/learn-go/go-todo/internal/ports"
)

// ------------------------------STRUCTS---------------------------------- //

type UserService struct {
	UserRepo ports.IUserRepositoryPort
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (s *UserService) Create(name string) (*models.User, error) {
	slug, err := s.UserRepo.GetSlug(name, nil)
	if err != nil {
		return nil, err
	}

	return s.UserRepo.Create(name, *slug)
}
