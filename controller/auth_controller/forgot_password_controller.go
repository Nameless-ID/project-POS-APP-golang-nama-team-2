package auth_controller

import (
	"context"
	"net/http"
	"project_pos_app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
	Token       string `json:"token" binding:"required"`
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi token dengan Redis
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

	// Hapus token dari Redis setelah digunakan
	err = h.RedisClient.Del(context.Background(), "reset_token:"+req.Email).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// validateResetToken memvalidasi token reset password menggunakan Redis
func (h *AuthHandler) validateResetToken(email, token string) (bool, error) {
	key := "reset_token:" + email
	storedToken, err := h.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil // Token tidak ditemukan
		}
		return false, err // Error Redis lainnya
	}

	// Bandingkan token yang dikirim dengan yang tersimpan
	return storedToken == token, nil
}

// updatePassword memperbarui password di database
func (h *AuthHandler) updatePassword(email, hashedPassword string) error {
	// Implementasikan logika update password di database
	return h.DB.Model(&User{}).Where("email = ?", email).Update("password", hashedPassword).Error
}
