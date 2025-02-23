package cache

import (
	"books-management-system/config"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache() *RedisCache {
	redisConfig := config.AppConfig.Redis
	redisAddr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully!")
	return &RedisCache{Client: client}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisCache) Set(ctx context.Context, key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisCache) DeleteMany(ctx context.Context, keys []string) error {
	return r.Client.Del(ctx, keys...).Err()
}

func (r *RedisCache) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.Client.Keys(ctx, pattern).Result()
}
