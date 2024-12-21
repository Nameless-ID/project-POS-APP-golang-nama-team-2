package repository

import (
	"errors"
	"project_pos_app/model"

	"gorm.io/gorm"
)

type Staff = model.Staff

type StaffRepository interface {
	GetAll() ([]Staff, error)
	GetByID(id int) (*Staff, error)
	Create(staff *Staff) error
	Update(staff *Staff) error
	Delete(id int) error
}

type staffRepository struct {
	DB *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{DB: db}
}

func (repo *staffRepository) GetAll() ([]Staff, error) {
	var staffs []Staff
	result := repo.DB.Find(&staffs)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffs, nil
}

func (repo *staffRepository) GetByID(id int) (*Staff, error) {
	var staff Staff
	result := repo.DB.First(&staff, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("staff not found")
	}
	return &staff, result.Error
}

func (repo *staffRepository) Create(staff *Staff) error {
	// Validasi usia sebelum disimpan
	if err := staff.ValidateAge(); err != nil {
		return err
	}

	result := repo.DB.Create(staff)
	return result.Error
}

func (repo *staffRepository) Update(staff *Staff) error {
	// Validasi usia sebelum diperbarui
	if err := staff.ValidateAge(); err != nil {
		return err
	}

	result := repo.DB.Save(staff)
	if result.RowsAffected == 0 {
		return errors.New("no record updated")
	}
	return result.Error
}

func (repo *staffRepository) Delete(id int) error {
	result := repo.DB.Delete(&Staff{}, id)
	if result.RowsAffected == 0 {
		return errors.New("staff not found or already deleted")
	}
	return result.Error
}
