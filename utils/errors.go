package utils

import "errors"

var (
	ErrBookNotFound  = errors.New("book not found")
	ErrBookDeletion  = errors.New("failed to delete book")
	ErrBookUpdate    = errors.New("failed to update book")
	ErrInvalidInput  = errors.New("invalid input data")
	ErrInvalidBookID = errors.New("invalid book ID")
	ErrInternalError = errors.New("internal server error")
)

type ErrorResponse struct {
	Message string `json:"message"`
}
