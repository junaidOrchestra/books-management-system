package cache

import "context"

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	Delete(ctx context.Context, key string) error
	DeleteMany(ctx context.Context, key []string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
}
