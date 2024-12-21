package service

import (
	"errors"
	"project_pos_app/model"
	repository "project_pos_app/repository/auth_repository"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo = repository.AuthRepoInterface

// AuthService mengelola logika bisnis autentikasi
type AuthService struct {
	Repo        *AuthRepo
	Log         *zap.Logger
	RedisClient *redis.Client
	DB          *gorm.DB
}

// NewAuthService membuat instance baru AuthService
func NewAuthService(repo *AuthRepo, log *zap.Logger, redisClient *redis.Client, db *gorm.DB) *AuthService {
	return &AuthService{
		Repo:        repo,
		Log:         log,
		RedisClient: redisClient,
		DB:          db,
	}
}

// ValidateUser memvalidasi email dan password user melalui repository
func (s *AuthService) ValidateUser(email, password string) (*model.User, error) {
	user, err := (*s.Repo).ValidateUser(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GenerateToken membuat JWT token untuk user
func (s *AuthService) GenerateToken(user *model.User) (string, error) {
	// Tentukan klaim untuk token
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix() // Token berlaku selama 24 jam

	// Tentukan secret key untuk signing token
	secretKey := "secret-key" // Sebaiknya menggunakan environment variable

	// Membuat token dengan signing method HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token yang sudah signed
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// SaveTokenToRedis menyimpan token ke Redis
func (s *AuthService) SaveTokenToRedis(user *model.User, token string) error {
	idKey := "session:" + user.Email
	// Simpan token ke Redis dengan expired time
	err := s.RedisClient.Set(s.RedisClient.Context(), idKey, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// ValidateToken memvalidasi JWT token dan mengembalikan klaimnya
func (s *AuthService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	secretKey := "secret-key" // Sebaiknya diambil dari environment variable

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan token menggunakan signing method yang benar
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Cek apakah token valid
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
