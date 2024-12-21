package repository

import (
	"fmt"
	"project_pos_app/model"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthRepoInterface adalah interface untuk repository autentikasi
type AuthRepoInterface interface {
	ValidateUser(email, password string) (*model.User, error)
}

// authRepo adalah implementasi dari AuthRepoInterface
type authRepo struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// NewAuthRepo membuat instance baru authRepo
func NewAuthRepo(db *gorm.DB, log *zap.Logger) AuthRepoInterface {
	return &authRepo{DB: db, Log: log}
}

// ValidateUser memvalidasi email dan password user
func (r *authRepo) ValidateUser(email, password string) (*model.User, error) {
	user := &model.User{}
	// Query user berdasarkan email
	if err := r.DB.Where("email = ?", email).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	// Bandingkan password yang di-hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}
