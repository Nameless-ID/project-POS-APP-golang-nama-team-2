package repository

import (
	"fmt"
	"net/smtp"
	"os"
	"project_pos_app/helper"
)

// EmailRepositoryInterface mendefinisikan kontrak untuk email repository
type otpRepository interface {
	SendOTPEmail(to, otp string) error
	
}

type OtpRepository struct{}

// NewEmailRepository membuat instance baru dari EmailRepository
func NewEmailRepository() *OtpRepository {
	return &OtpRepository{}
}

// SendOTPEmail mengirimkan OTP melalui email
func (r *OtpRepository) SendOTPEmail(to, otp string) error {
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
		helper.LogError("Failed to send OTP email", err) // Log jika terjadi error
		return err
	}
	return nil
}
