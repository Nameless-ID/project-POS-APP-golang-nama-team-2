package service

import (
	"errors"
	"project_pos_app/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (a *AuthService) ValidateUser(email, password string) (*model.User, error) {
	user := &model.User{}
	// Query user berdasarkan email
	if err := a.DB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Bandingkan password yang di-hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
