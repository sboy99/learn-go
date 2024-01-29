package repositories

import "github.com/sboy99/learn-go/go-task-management/src/internal/domain/models"

type IUserRepository interface {
	Create(name string, email string) *models.User
}
