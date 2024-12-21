package controller

import (
	"net/http"
	"project_pos_app/database"
	"project_pos_app/helper"
	"project_pos_app/model"
	"project_pos_app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service     *service.AllService
	Log         *zap.Logger
	Cacher      *database.Cache
	RedisClient *redis.Client
	DB          *gorm.DB
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(service *service.AllService, log *zap.Logger, rdb *database.Cache, redisClient *redis.Client, db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		Service:     service,
		Log:         log,
		RedisClient: redisClient,
		Cacher:      rdb,
		DB:          db,
	}
}

// Login handles the login request
func (auth *AuthHandler) Login(c *gin.Context) {
	// 1. Initialize the Login struct
	login := model.Login{}

	// 2. Bind JSON request to struct
	if err := c.ShouldBindJSON(&login); err != nil {
		auth.Log.Error("Invalid payload", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "Invalid Payload: "+err.Error(), nil)
		return
	}

	// 3. Validate user credentials using service layer
	user, err := auth.Service.Auth.ValidateUser(login.Email, login.Password)
	if err != nil {
		auth.Log.Error("Failed to validate user", zap.Error(err))
		helper.Responses(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	// 4. Generate JWT Token
	secretKey := "secret-key" // Should be retrieved from environment variables for security
	token, err := helper.GenerateToken(uint(user.ID), user.Email, user.Role, secretKey)
	if err != nil {
		auth.Log.Error("Failed to generate token", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	// 5. Store token in Redis with unique key (example: IDKEY)
	idKey := "session:" + user.Email
	auth.Log.Info("Saving token to Redis", zap.String("IDKEY", idKey), zap.String("token", token))

	if err := auth.Cacher.Set(idKey, token); err != nil {
		auth.Log.Error("Failed to save token in Redis", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "Failed to save session", nil)
		return
	}

	// 6. Successful response
	session := map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
			"role":  user.Role,
		},
	}

	helper.Responses(c, http.StatusOK, "Successfully logged in", session)
}
