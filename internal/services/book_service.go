package services

import (
	"books-management-system/internal/models"
	"books-management-system/internal/repositories"
	"books-management-system/pkg/cache"
	"books-management-system/pkg/kafka"
	"books-management-system/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BookService struct {
	Repo     repositories.BookRepository
	Cache    cache.Cache
	Producer *kafka.Producer
}

func NewBookService(repo repositories.BookRepository, cache cache.Cache, producer *kafka.Producer) *BookService {
	return &BookService{Repo: repo, Cache: cache, Producer: producer}
}

func (s *BookService) GetBooks(ctx context.Context, page, limit int) ([]models.Book, error) {
	cacheKey := utils.BooksPageKey(page, limit)

	// Try fetching from cache
	if s.Cache != nil {
		cachedData, err := s.Cache.Get(ctx, cacheKey)
		if err == nil {
			var books []models.Book
			if json.Unmarshal([]byte(cachedData), &books) == nil {
				return books, nil // ✅ Cache hit
			}
		} else if err != redis.Nil {
			utils.Logger.Warnw("Redis error while fetching books", "error", err)
		}
	}

	// Fetch from database
	books, err := s.Repo.GetBooks(page, limit)
	if err != nil {
		utils.Logger.Errorw("Database error while fetching books", "error", err)
		return nil, utils.ErrInternalError
	}

	s.cacheDataAsync(ctx, cacheKey, books)

	return books, nil
}

func (s *BookService) GetBookByID(ctx context.Context, id uint) (*models.Book, error) {
	cacheKey := utils.BookKey(id)

	if s.Cache != nil {
		cachedData, err := s.Cache.Get(ctx, cacheKey)
		if err == nil {
			var book models.Book
			if json.Unmarshal([]byte(cachedData), &book) == nil {
				return &book, nil
			}
		} else if err != redis.Nil {
			utils.Logger.Warn("Redis error while fetching book", err)
		}
	}

	book, err := s.Repo.GetBookByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Logger.Infow("Book not found", "book_id", id)
			return nil, utils.ErrBookNotFound
		}
		utils.Logger.Error("Database error while fetching book", err)
		return nil, utils.ErrInternalError
	}

	s.cacheDataAsync(ctx, cacheKey, book)

	return book, nil
}

func (s *BookService) CreateBook(ctx context.Context, book *models.Book) error {
	if err := s.Repo.CreateBook(book); err != nil {
		utils.Logger.Error("Failed to create book:", err)
		return utils.ErrInternalError
	}

	go func() {
		if err := s.Producer.Publish(kafka.TopicBookEvents, kafka.EventBookCreated, book); err != nil {
			utils.Logger.Error("Failed to publish book creation event:", err)
		}
	}()
	return nil
}

func (s *BookService) UpdateBook(ctx context.Context, book *models.Book) error {
	if err := s.Repo.UpdateBook(book); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrBookNotFound
		}
		utils.Logger.Error("Failed to update book:", err)
		return utils.ErrInternalError
	}

	if s.Cache != nil {
		if err := s.Cache.Delete(ctx, utils.BookKey(book.ID)); err != nil {
			utils.Logger.Error("Failed to delete book from cache:", err)
		}
		s.invalidatePaginatedCache(ctx)
	}

	go func() {
		if err := s.Producer.Publish(kafka.TopicBookEvents, kafka.EventBookUpdated, book); err != nil {
			utils.Logger.Error("Failed to publish book update event:", err)
		}
	}()

	return nil
}

func (s *BookService) DeleteBook(ctx context.Context, id uint) error {
	err := s.Repo.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrBookNotFound
		}
		return utils.ErrInternalError
	}

	if s.Cache != nil {
		if err := s.Cache.Delete(ctx, utils.BookKey(id)); err != nil {
			utils.Logger.Error("Failed to delete book from cache:", err)
		}
		s.invalidatePaginatedCache(ctx)
	}

	go func() {
		if err := s.Producer.Publish(kafka.TopicBookEvents, kafka.EventBookDeleted, id); err != nil {
			utils.Logger.Error("Failed to publish book deletion event:", err)
		}
	}()

	return nil
}

func (s *BookService) cacheDataAsync(ctx context.Context, key string, data interface{}) {
	if s.Cache == nil {
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		utils.Logger.Warnw("Failed to serialize data", "cache_key", key, "error", err)
		return
	}

	go func() {
		if err := s.Cache.Set(ctx, key, string(jsonData)); err != nil {
			utils.Logger.Warnw("Failed to cache data", "cache_key", key, "error", err)
		}
	}()
}

func (s *BookService) invalidatePaginatedCache(ctx context.Context) {
	if s.Cache == nil {
		return
	}

	keys, err := s.Cache.Keys(ctx, "books:page_*")
	if err != nil {
		utils.Logger.Error("Failed to fetch cache keys:", err)
		return
	}

	if len(keys) > 0 {
		if err := s.Cache.DeleteMany(ctx, keys); err != nil {
			utils.Logger.Error("Failed to delete cache keys:", err)
		}
	}
}
