package services

import (
	"github.com/sboy99/learn-go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/learn-go/go-task-management/src/internal/repositories"
)

// user service structures
type UserService struct {
	UserRepo repositories.IUserRepository
}

// user service methods
func (s *UserService) Create(name string, email string) *models.User {
	return s.UserRepo.Create(name, email)
}

func (s *UserService) List() *[]models.User {
	return nil
}

func (s *UserService) GetById(id int64) *models.User {
	return nil
}

func (s *UserService) UpdateById(id int64) *models.User {
	return nil
}

func (s *UserService) DeleteById(id int64) *models.User {
	return nil
}
