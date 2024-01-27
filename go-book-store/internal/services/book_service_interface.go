package services

import (
	"github.com/sboy99/learn-go/go-book-store/internal/domain/models"
)

type IBookService interface {
	CreateBook(name string, price float64) *models.Book
	ListBooks() *[]models.Book
	GetOneBook(id int64) *models.Book
}
