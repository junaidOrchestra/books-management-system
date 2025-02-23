package sqlite

import (
	"books-management-system/internal/models"
	"books-management-system/internal/repositories"
	"gorm.io/gorm"
)

type SQLiteBookRepository struct {
	DB *gorm.DB
}

// NewSQLiteBookRepository returns an implementation of BookRepository
func NewSQLiteBookRepository(db *gorm.DB) repositories.BookRepository {
	db.AutoMigrate(&models.Book{})
	return &SQLiteBookRepository{DB: db}
}

func (r *SQLiteBookRepository) GetBooks(page, limit int) ([]models.Book, error) {
	var books []models.Book
	offset := (page - 1) * limit
	err := r.DB.Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *SQLiteBookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	result := r.DB.First(&book, id)
	return &book, result.Error
}

func (r *SQLiteBookRepository) CreateBook(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *SQLiteBookRepository) UpdateBook(book *models.Book) error {
	return r.DB.Save(book).Error
}

func (r *SQLiteBookRepository) DeleteBook(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}
