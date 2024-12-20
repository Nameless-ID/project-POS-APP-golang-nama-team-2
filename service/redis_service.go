package service

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Client *redis.Client
}

// NewRedisService membuat instance RedisService baru
func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{Client: client}
}

// Set menyimpan key-value pair ke Redis dengan masa berlaku tertentu
func (r *RedisService) Set(key, value string, expiration time.Duration) error {
	ctx := context.Background()
	err := r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get mengambil value berdasarkan key dari Redis
func (r *RedisService) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil // Key tidak ditemukan
		}
		return "", err // Error Redis lainnya
	}
	return value, nil
}

// Delete menghapus key dari Redis
func (r *RedisService) Delete(key string) error {
	ctx := context.Background()
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
