package controller

import (
	"context"
	"net/http"
	"project_pos_app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
	Token       string `json:"token" binding:"required"`
}

// ResetPassword menangani permintaan reset password
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest

	// Validasi input JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi token menggunakan Redis
	isValid, err := h.validateResetToken(req.Email, req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating token"})
		return
	}

	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Hash password baru
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update password di database
	err = h.updatePassword(req.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	// Hapus token dari Redis setelah berhasil
	err = h.RedisClient.Del(context.Background(), "reset_token:"+req.Email).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// validateResetToken memvalidasi token reset password menggunakan Redis
func (h *AuthHandler) validateResetToken(email, token string) (bool, error) {
	ctx := context.Background()
	key := "reset_token:" + email

	// Ambil token dari Redis
	storedToken, err := h.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil // Token tidak ditemukan
		}
		return false, err // Error Redis lainnya
	}

	// Periksa kecocokan token
	return storedToken == token, nil
}

// updatePassword memperbarui password di database
func (h *AuthHandler) updatePassword(email, hashedPassword string) error {
	// Update password dengan GORM
	result := h.DB.Model(&User{}).Where("email = ?", email).Update("password", hashedPassword)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound // Email tidak ditemukan
	}
	return nil
}

// User adalah representasi tabel pengguna di database
type User struct {
	Email    string `gorm:"primaryKey"`
	Password string
}
