package cache

import (
	"context"
	"errors"
	"main/services/api-gateway/internal/config"
	"main/services/api-gateway/internal/values"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, key string) error
}

type redisClient struct {
	Client *redis.Client
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *redisClient) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if exists == 0 {
		return false, values.ErrRedisKeyExists
	}
	return true, nil
}

func (r *redisClient) Delete(ctx context.Context, key string) error {
	_, err := r.Client.Del(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}

func NewCache(cfg config.RedisConfig) (Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &redisClient{
		Client: client,
	}, nil
}
