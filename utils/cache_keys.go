package utils

import "fmt"

//type CacheKeys struct{}

//func (c CacheKeys) BookKey(id uint) string {
//	return fmt.Sprintf("book:%d", id) // ✅ Generates book-specific cache key
//}
//
//func (c CacheKeys) BooksPageKey(page, limit int) string {
//	return fmt.Sprintf("books:page_%d_limit_%d", page, limit) // ✅ Key for paginated books
//}

func BookKey(id uint) string {
	return fmt.Sprintf("book:%d", id) // ✅ Generates book-specific cache key
}

func BooksPageKey(page, limit int) string {
	return fmt.Sprintf("books:page_%d_limit_%d", page, limit) // ✅ Key for paginated books
}
