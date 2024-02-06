package services

import (
	"errors"

	"github.com/sboy99/learn-go/go-todo/infra/exception"
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
	"github.com/sboy99/learn-go/go-todo/internal/helpers"
	"github.com/sboy99/learn-go/go-todo/internal/ports"
)

// ------------------------------STRUCTS---------------------------------- //

type AuthService struct {
	UserRepo ports.IUserRepositoryPort
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (s *AuthService) Register(email string, pass string, name string) (*models.User, *exception.HttpException) {
	// check if user already registered
	_, err := s.checkForExistingEmail(email)
	if err != nil {
		return nil, exception.BadRequestException(err)
	}

	// encrypt password
	hashedPass, err := s.encryptPassword(pass)
	if err != nil {
		return nil, exception.BadRequestException(err)
	}

	// generate unique slug/username
	slug, err := s.UserRepo.GetSlug(name)
	if err != nil {
		return nil, exception.BadRequestException(err)
	}

	// create new user
	user, err := s.UserRepo.Create(email, hashedPass, name, slug)
	if err != nil {
		return nil, exception.InternalServerErrorException(err)
	}

	// done
	return user, nil
}

// ------------------------------PRIVATE_METHODS---------------------------------- //

func (s *AuthService) checkForExistingEmail(email string) (bool, error) {
	user, _ := s.UserRepo.GetUserWithEmail(email)

	if user == nil {
		return false, nil
	}

	return true, errors.New(`Email already exists`)
}

func (s *AuthService) encryptPassword(password string) (string, error) {
	return helpers.HashStr(password)
}
