package repository

import (
	"fmt"
	"project_pos_app/model"

	"gorm.io/gorm"
)

// EmailRepoInterface mendefinisikan kontrak untuk repositori email
type emailRepo interface {
	GetPasswordByEmail(email string) (string, error)
}

// EmailRepo adalah implementasi dari EmailRepoInterface
type EmailRepo struct {
	DB *gorm.DB
}

// NewEmailRepo membuat instance baru EmailRepo
func NewEmailRepo() *EmailRepo {
	return &EmailRepo{}
}

// GetPasswordByEmail mendapatkan password pengguna berdasarkan email
func (r *EmailRepo) GetPasswordByEmail(email string) (string, error) {
	var user model.User

	// Query menggunakan GORM untuk mencari pengguna berdasarkan email
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Jika pengguna dengan email tersebut tidak ditemukan
			return "", fmt.Errorf("email not found")
		}
		// Jika terjadi error lainnya
		return "", fmt.Errorf("error fetching user: %v", err)
	}

	// Kembalikan password pengguna jika ditemukan
	return user.Password, nil
}
