package helper

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims adalah struktur klaim yang terdapat di dalam JWT
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.SigningMethodHMAC
}

// AuthMiddleware adalah middleware untuk validasi JWT
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Hapus prefix "Bearer " jika ada
		token = strings.TrimPrefix(token, "Bearer ")

		// Validasi token
		claims, err := ValidateToken(token, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set klaim ke context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

// // ValidateToken untuk memvalidasi JWT dan mengembalikan klaim
// func ValidateToken(tokenString, secretKey string) (*Claims, error) {
// 	// Parse token dan klaim
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		// Pastikan token menggunakan signing method yang benar
// 		if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	// Cek apakah token valid
// 	claims, ok := token.Claims.(*Claims)
// 	if !ok || !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	// Kembalikan klaim yang valid
// 	return claims, nil
// }
