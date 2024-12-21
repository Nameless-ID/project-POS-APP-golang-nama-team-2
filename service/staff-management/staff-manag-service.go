package service

import (
	"errors"
	"project_pos_app/model"
	repository "project_pos_app/repository/staff_management"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Staff = model.Staff
type StaffRepo = repository.StaffRepository

type staffService struct {
	Repo StaffRepo
	DB   *gorm.DB
	Log  *zap.Logger
}

type StaffService interface {
	GetAll() ([]Staff, error)
	GetByID(id int) (*Staff, error)
	Create(staff *Staff) error
	Update(staff *Staff) error
	Delete(id int) error
}

func NewStaffService(repo StaffRepo, db *gorm.DB, log *zap.Logger) StaffService {
	return &staffService{
		Repo: repo,
		DB:   db,
		Log:  log,
	}
}

// GetAll retrieves all staff records
func (s *staffService) GetAll() ([]Staff, error) {
	staffs, err := s.Repo.GetAll()
	if err != nil {
		s.Log.Error("Failed to retrieve all staff", zap.Error(err))
		return nil, err
	}
	return staffs, nil
}

// GetByID retrieves a staff by ID
func (s *staffService) GetByID(id int) (*Staff, error) {
	staff, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.Warn("Staff not found", zap.Int("ID", id))
		} else {
			s.Log.Error("Failed to retrieve staff by ID", zap.Int("ID", id), zap.Error(err))
		}
		return nil, err
	}
	return staff, nil
}

// Create adds a new staff record
func (s *staffService) Create(staff *Staff) error {
	err := s.Repo.Create(staff)
	if err != nil {
		s.Log.Error("Failed to create staff", zap.Error(err))
		return err
	}
	s.Log.Info("Staff created successfully", zap.Int("ID", staff.ID))
	return nil
}

// Update modifies an existing staff record
func (s *staffService) Update(staff *Staff) error {
	err := s.Repo.Update(staff)
	if err != nil {
		s.Log.Error("Failed to update staff", zap.Int("ID", staff.ID), zap.Error(err))
		return err
	}
	s.Log.Info("Staff updated successfully", zap.Int("ID", staff.ID))
	return nil
}

// Delete removes a staff record by ID
func (s *staffService) Delete(id int) error {
	err := s.Repo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Log.Warn("Staff not found for deletion", zap.Int("ID", id))
		} else {
			s.Log.Error("Failed to delete staff", zap.Int("ID", id), zap.Error(err))
		}
		return err
	}
	s.Log.Info("Staff deleted successfully", zap.Int("ID", id))
	return nil
}
