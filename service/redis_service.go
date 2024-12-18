package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type RedisService struct {
	Client *redis.Client
}

func NewRedisService(addr string) *RedisService {
	client := redis.NewClient(&redis.Options{Addr: addr})
	return &RedisService{Client: client}
}

func (r *RedisService) Set(key, value string, expiration time.Time) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisService) Get(key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
