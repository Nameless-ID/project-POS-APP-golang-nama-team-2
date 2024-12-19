package controllers

import (
	"auth_jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateOTP(redisService *services.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		otp := c.Query("otp")

		savedOTP, err := redisService.Get(email)
		if err != nil || savedOTP != otp {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid OTP"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OTP is valid"})
	}
}
