package model

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"project_pos_app/utils"
	"time"

	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string          `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Email     string          `gorm:"type:varchar(255);unique" json:"email" binding:"required,email"`
	Password  string          `gorm:"type:varchar(255)" json:"password" binding:"required,min=8"`
	Role      string          `gorm:"type:varchar(255)" json:"role" binding:"required"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type Session struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"type:int"`
	Token        string    `gorm:"not null"`
	IpAddress    string    `gorm:"not null"`
	LastActivity time.Time `gorm:"not null"`
}

// UserSeed generates initial user data for seeding (without database operations)
func UserSeed() *User {
	// Generate RSA Key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate RSA key: %v", err)
	}

	password := "superadmin123" // Default password for Super Admin

	// Hash the password
	hashedPassword, err := utils.GeneratePassword(password, privateKey)
	if err != nil {
		log.Fatalf("Error hashing password for seeder: %v", err)
	}

	return &User{
		Name:     "Super Admin",
		Email:    "superadmin@example.com",
		Password: hashedPassword,
		Role:     "super_admin",
	}
}

// SeedSuperAdmin seeds the Super Admin data into the database
func SeedSuperAdmin(db *gorm.DB) {
	// Get Super Admin user data
	superAdmin := UserSeed()

	// Check or Insert Super Admin into the database
	if err := db.FirstOrCreate(&superAdmin, User{Email: superAdmin.Email}).Error; err != nil {
		fmt.Printf("Failed to seed Super Admin: %v\n", err)
	} else {
		fmt.Println("Super Admin seeded successfully")
	}
}
