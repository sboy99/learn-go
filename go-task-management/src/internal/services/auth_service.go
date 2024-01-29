package services

import (
	"github.com/sboy99/learn-go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/learn-go/go-task-management/src/internal/helpers"
	"github.com/sboy99/learn-go/go-task-management/src/internal/repositories"
)

type AuthService struct {
	UserRepo     repositories.IUserRepository
	CryptoHelper helpers.ICryptoHelper
}

func (s *AuthService) Register(name string, email string, password string) *models.User {
	// check for existing emails
	s._CheckForRegisteredEmail(email)
	// create user
	user := s._CreateUser(name, email, password)
	// return user
	return user
}

func (s *AuthService) _CheckForRegisteredEmail(email string) {
	user := s.UserRepo.GetUserByEmail(email)

	if user != nil {
		// todo: throw http exception
		return
	}
}

func (s *AuthService) _CreateUser(name string, email string, password string) *models.User {
	// get user slug
	userSlug := s.UserRepo.CreateSlug(name)
	// get hash password
	hashedPassword := s.CryptoHelper.GetHash(password)
	// create a new user
	return s.UserRepo.Create(userSlug, name, email, hashedPassword)
}
