package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	service "project_pos_app/service/auth_service"
)

type EmailHandler struct {
	Service *service.EmailService
	Log     *zap.Logger
}

// NewAuthHandler membuat instance AuthHandler
func NewEmailHandler(service *service.EmailService, log *zap.Logger) *AuthHandler {
	return &AuthHandler{
		Service: service,
		Log:     log,
	}
}

// CheckEmail memeriksa apakah email valid dan mengirimkan password jika valid
func (h *AuthHandler) CheckEmail(c *gin.Context) {
	// Ambil email dari query parameter
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Kirim email dengan password
	err := h.Service.SendPasswordEmail(email)
	if err != nil {
		h.Log.Error("Failed to send email", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password has been sent to your email"})
}
