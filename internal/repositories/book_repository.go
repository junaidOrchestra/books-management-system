package repositories

import "books-management-system/internal/models"

type BookRepository interface {
	GetBooks(page, limit int) ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}
