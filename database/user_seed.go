package database

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"project_pos_app/helper"
	"project_pos_app/model"
)

// UserSeed generates initial user data for seeding (without database operations)
func UserSeed() *model.User {
	// Generate RSA Key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate RSA key: %v", err)
	}

	password := "superadmin123" // Default password for Super Admin

	// Hash the password
	hashedPassword, err := helper.GeneratePassword(password, privateKey)
	if err != nil {
		log.Fatalf("Error hashing password for seeder: %v", err)
	}

	return &model.User{
		Name:     "Super Admin",
		Email:    "superadmin@example.com",
		Password: hashedPassword,
		Role:     "super_admin",
	}
}
