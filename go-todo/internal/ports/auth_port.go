package ports

import (
	"github.com/sboy99/learn-go/go-todo/infra/exception"
	"github.com/sboy99/learn-go/go-todo/internal/domain/models"
)

// -------------------------------INTERFACES--------------------------------- //

type IAuthServicePort interface {
	Register(email string, pass string, name string) (*models.User, *exception.HttpException)
	Login(email string, pass string) (*models.User, *exception.HttpException)
}
