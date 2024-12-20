package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Staff adalah representasi dari tabel "staff" dalam database
type Staff struct {
	ID        int             `gorm:"primaryKey" json:"id"`                              // ID unik sebagai primary key
	Email     string          `gorm:"uniqueIndex" json:"email" binding:"required,email"` // Email unik dan wajib
	Password  string          `json:"-" binding:"required,min=8"`                        // Password dengan minimal 8 karakter
	Role      string          `json:"role" binding:"required"`                           // Peran (contoh: admin, user, dsb.)
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`                  // Waktu pembuatan otomatis
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`                  // Waktu pembaruan otomatis
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`                                    // Soft delete untuk mengarsipkan data

	Photo            []byte    `json:"photo,omitempty"`                  // Foto dalam bentuk byte array
	BirthOfDate      time.Time `json:"birth_of_date" binding:"required"` // Tanggal lahir wajib diisi
	Age              int       `json:"age" binding:"gte=0,lte=150"`      // Usia dengan rentang validasi 0-150
	Salary           float64   `json:"salary" binding:"gte=0"`           // Gaji harus bernilai positif
	Timing           string    `json:"timing"`                           // Waktu shift atau lainnya
	AdditionalDetail string    `json:"additional_detail,omitempty"`      // Detail tambahan opsional
}

// BeforeCreate adalah hook yang dijalankan sebelum pembuatan data baru
func (s *Staff) BeforeCreate(tx *gorm.DB) (err error) {
	// Contoh: Otomatis menghitung usia berdasarkan tanggal lahir
	if !s.BirthOfDate.IsZero() {
		currentYear := time.Now().Year()
		birthYear := s.BirthOfDate.Year()
		s.Age = currentYear - birthYear
	}

	// Tambahkan logika lainnya jika perlu
	return
}

// ValidateAge memvalidasi apakah usia cocok dengan tanggal lahir
func (s *Staff) ValidateAge() error {
	expectedAge := time.Now().Year() - s.BirthOfDate.Year()
	if s.Age != expectedAge {
		return errors.New("age does not match the date of birth")
	}
	return nil
}
