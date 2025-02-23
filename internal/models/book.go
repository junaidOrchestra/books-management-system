package models

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"gt=500"`
}
