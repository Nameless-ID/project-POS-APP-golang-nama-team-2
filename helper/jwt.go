package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims defines the custom claims structure for JWT
type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token with custom claims
func GenerateToken(userID uint, email, role, secretKey string) (string, error) {
	// Custom claims
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-auth-gin",                                     // Issuer
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Issued at time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expiration time (24 hours)
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString, secretKey string) (*CustomClaims, error) {
	// Parse the token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and return the claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
