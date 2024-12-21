package service

import (
	"fmt"
	"net/smtp"
	repository "project_pos_app/repository/auth_repository"
)

type Email = repository.EmailRepo

// type emailService interface {
// }

// EmailService mengelola logika bisnis untuk pengiriman email
type EmailService struct {
	Repo Email
}

// NewEmailService membuat instance baru EmailService
func NewEmailService(repo repository.EmailRepo) *EmailService {
	return &EmailService{Repo: repo}
}

// SendPasswordEmail mengirimkan email berisi password ke pengguna
func (s *EmailService) SendPasswordEmail(email string) error {
	// Cari password berdasarkan email
	password, err := s.Repo.GetPasswordByEmail(email)
	if err != nil {
		return fmt.Errorf("email not found: %v", err)
	}

	// Kirim email dengan password
	err = sendEmail(email, password)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

// sendEmail mengirim email dengan SMTP
func sendEmail(to string, password string) error {
	// Konfigurasi email
	from := "your-email@example.com"       // Ganti dengan email pengirim
	passwordEmail := "your-email-password" // Ganti dengan password email pengirim
	smtpHost := "smtp.example.com"         // Ganti dengan host SMTP (misal: smtp.gmail.com)
	smtpPort := "587"                      // Ganti dengan port SMTP (misal: 587 untuk TLS)

	// Isi email
	subject := "Your Password"
	body := fmt.Sprintf("Hello,\n\nYour password is: %s\n\nRegards,\nYour Team", password)
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// Kirim email menggunakan SMTP
	auth := smtp.PlainAuth("", from, passwordEmail, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
