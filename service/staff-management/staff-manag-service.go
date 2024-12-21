package service

import (
	"project_pos_app/model"
	repository "project_pos_app/repository/staff_management"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Staff = model.Staff
type staffRepo = *repository.StaffRepository

type staffService struct {
	Repo staffRepo
	DB   *gorm.DB
	Log  *zap.Logger
}

type StaffService interface {
	GetAll() []Staff
	GetByID(id int) *Staff
	Create(staff *Staff) error
	Update(staff *Staff) error
	Delete(id int) error
}

func NewStaffService(repo staffRepo, db *gorm.DB, log *zap.Logger) StaffService {
	return &staffRepository{
		DB:   db,
		Repo: repo,
		Log:  log,
	}
}

func (s *staffService) GetAll() []Staff {
	return
}

func (s *staffService) GetByID() []Staff {
	return
}

func (s *staffService) Create() []Staff {
	return
}
func (s *staffService) Update() []Staff {
	return
}
func (s *staffService) Delete() []Staff {
	return
}
