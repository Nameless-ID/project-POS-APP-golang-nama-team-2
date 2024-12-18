package model

import "time"

type OTP struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"index" json:"email"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
}
