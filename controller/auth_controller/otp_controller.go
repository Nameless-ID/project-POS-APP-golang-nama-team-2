package controller

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// Load .env file
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

// Kirim OTP ke email
func sendOTPEmail(to, otp string) error {
	// Ambil konfigurasi SMTP dari environment variables
	from := os.Getenv("SMTP_EMAIL")
	passwordEmail := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Subjek dan body email
	subject := "Your OTP Code"
	body := fmt.Sprintf("Hello,\n\nYour OTP code is: %s\n\nRegards,\nYour Team", otp)
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// Kirim email
	auth := smtp.PlainAuth("", from, passwordEmail, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
