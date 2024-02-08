package services

import (
	"errors"

	"github.com/sboy99/learn-go/go-todo/infra/exception"
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
	"github.com/sboy99/learn-go/go-todo/internal/ports"
)

// ------------------------------STRUCTS---------------------------------- //

type AuthService struct {
	UserRepo     ports.IUserRepositoryPort
	JwtAdapter   ports.IJWTAdapterPort
	CryptoHelper ports.ICryptoHelperPort
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (s *AuthService) Register(email string, pass string, name string) (*models.User, *exception.HttpException) {
	// check if user already registered
	_, err := s.isExistingEmail(email)
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

func (s *AuthService) Login(email string, pass string) (string, string, *exception.HttpException) {
	// get user
	user, err := s.getUserByEmail(email)
	if user == nil || err != nil {
		return "", "", exception.NotFoundException(err)
	}

	// compare password
	if err := s.comparePassword(pass, *&user.Pass); err != nil {
		return "", "", exception.BadRequestException(err)
	}

	// user jwt payload
	payload := s.getUserJwtPayload(user)

	// get jwt token
	accessToekn, err := s.JwtAdapter.Sign(payload)
	if err != nil {
		return "", "", exception.InternalServerErrorException(err)
	}

	// done
	return accessToekn, "", nil
}

// ------------------------------PRIVATE_METHODS---------------------------------- //

func (s *AuthService) isExistingEmail(email string) (bool, error) {
	user, _ := s.UserRepo.GetUserByEmail(email)

	if user == nil {
		return false, nil
	}

	return true, errors.New(`Email already exists`)
}

func (s *AuthService) encryptPassword(password string) (string, error) {
	return s.CryptoHelper.HashStr(password)
}

func (s *AuthService) getUserByEmail(email string) (*models.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *AuthService) comparePassword(password string, hash string) error {
	isCorrect := s.CryptoHelper.CompareHash(password, hash)

	if !isCorrect {
		return errors.New(`Invalid Credentials`)
	}

	return nil
}

func (s *AuthService) getUserJwtPayload(user *models.User) map[string]interface{} {
	payload := make(map[string]interface{})

	payload["sub"] = user.Slug
	payload["id"] = user.Id
	payload["name"] = user.Name

	return payload
}
