package config

import (
	"github.com/go-redis/redis/v8"
)

// NewRedisClient menginisialisasi koneksi Redis
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Ganti dengan alamat Redis Anda
		Password: "",               // Ganti jika Redis memerlukan password
		DB:       0,                // Gunakan database default (0)
	})
}
