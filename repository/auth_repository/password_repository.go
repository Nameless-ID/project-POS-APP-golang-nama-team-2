package repository

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type PasswordRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func (pr *PasswordRepository) SendPasswordResetEmail(email string) error {
	// Memuat konfigurasi dari file .env
	err := godotenv.Load()
	if err != nil {
		pr.Log.Error("Error loading .env file", zap.Error(err))
		return err
	}

	// Membaca konfigurasi dari environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	// smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")

	// Membuat pesan email
	subject := "Password Reset Request"
	body := fmt.Sprintf(
		"Hello,\n\nPlease reset your password by clicking the link below:\nhttp://example.com/reset-password?email=%s\n\nThanks!",
		email,
	)

	// Membuat pesan email menggunakan Gomail
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Konfigurasi Dialer untuk koneksi SMTP
	d := gomail.NewDialer(smtpHost, 587, senderEmail, senderPassword)

	// Mengirim email
	if err := d.DialAndSend(m); err != nil {
		pr.Log.Error("Failed to send password reset email", zap.Error(err))
		return err
	}

	pr.Log.Info("Password reset email sent successfully", zap.String("email", email))
	return nil
}
