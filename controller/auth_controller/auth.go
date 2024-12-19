package auth_controller

import (
	"net/http"
	"project_pos_app/database"
	"project_pos_app/helper"
	"project_pos_app/model"
	"project_pos_app/service"
	"project_pos_app/utils" // Import utils untuk JWT

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service *service.AllService
	Log     *zap.Logger
	Cacher  *database.Cache
}

// NewAuthHandler membuat instance AuthHandler
func NewAuthHandler(service *service.AllService, log *zap.Logger, rdb *database.Cache) *AuthHandler {
	return &AuthHandler{
		Service: service,
		Log:     log,
		Cacher:  rdb,
	}
}

// Login menangani permintaan login
func (auth *AuthHandler) Login(c *gin.Context) {
	// 1. Inisialisasi struct Login
	login := model.Login{}

	// 2. Bind JSON request ke struct
	if err := c.ShouldBindJSON(&login); err != nil {
		auth.Log.Error("Invalid payload", zap.Error(err))
		helper.Responses(c, http.StatusBadRequest, "Invalid Payload: "+err.Error(), nil)
		return
	}

	// 3. Validasi user di service layer
	user, err := auth.Service.Auth.ValidateUser(login.Email, login.Password)
	if err != nil {
		auth.Log.Error("Failed to validate user", zap.Error(err))
		helper.Responses(c, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	// 4. Generate JWT Token
	secretKey := "secret-key" // Ambil dari environment untuk keamanan
	token, err := utils.GenerateToken(uint(user.ID), user.Email, user.Role, secretKey)
	if err != nil {
		auth.Log.Error("Failed to generate token", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	// 5. Simpan token ke Redis dengan kunci unik (contoh: IDKEY)
	idKey := "session:" + user.Email
	auth.Log.Info("Saving token to Redis", zap.String("IDKEY", idKey), zap.String("token", token))

	if err := auth.Cacher.Set(idKey, token); err != nil {
		auth.Log.Error("Failed to save token in Redis", zap.Error(err))
		helper.Responses(c, http.StatusInternalServerError, "Failed to save session", nil)
		return
	}

	// 6. Respon sukses
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
