package model

import (
	"crypto/rand"
	"crypto/rsa"
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
	Name      string          `grom:"type:varchar(100)" json:"name" binding:"required"`
	Email     string          `grom:"type:varchar(255);unique" json:"email" binding:"required,email"`
	Password  string          `grom:"type:varchar(50)" json:"password" binding:"required,min=8"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
	Role      string          `gorm:"type:varchar(255)" json:"role" binding:"required"`
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type Session struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int       `gorm:"type:int"`
	Token        string    `gorm:"not null"`
	IpAddress    string    `gorm:"not null"`
	LastActivity time.Time `gorm:"not null"`
}

// UserSeed generates initial user data for Super Admin
func UserSeed() interface{} {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048) // Generate RSA Key
	password := "superadmin123"                         // Default password for Super Admin

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
