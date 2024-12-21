package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type otpService struct {
	repo   *otpRepository
	client *redis.Client
	logger *zap.Logger
}

func NewOtpService(client *redis.Client, logger *zap.Logger) *otpService {
	return &otpService{
		client: client,
		logger: logger,
	}
}

// GenerateOTP untuk menghasilkan OTP 6 digit
func GenerateOTP() string {
	max := new(big.Int)
	max.Exp(big.NewInt(10), big.NewInt(6), nil)
	otp, _ := rand.Int(rand.Reader, max)
	return fmt.Sprintf("%06d", otp)
}

// Set OTP di Redis
func (r *otpService) Set(email, otp string) error {
	// Menyimpan OTP dengan waktu kadaluarsa 5 menit (300 detik)
	err := r.client.Set(r.client.Context(), email, otp, time.Minute*5).Err()
	if err != nil {
		r.logger.Error("failed to set OTP in Redis", zap.Error(err))
		return err
	}
	return nil
}

// Get OTP dari Redis
func (r *otpService) Get(email string) (string, error) {
	otp, err := r.client.Get(r.client.Context(), email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("OTP not found")
		}
		r.logger.Error("failed to get OTP from Redis", zap.Error(err))
		return "", err
	}
	return otp, nil
}
