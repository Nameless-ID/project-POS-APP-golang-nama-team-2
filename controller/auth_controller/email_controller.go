package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) CheckEmail(c *gin.Context) {
	// Ambil email dari query parameter
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Logika validasi email
	// Misalnya, cek apakah email ada di database
	isValid := true // Ganti dengan logika sebenarnya

	if isValid {
		c.JSON(http.StatusOK, gin.H{"message": "Email is valid"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
	}
}
