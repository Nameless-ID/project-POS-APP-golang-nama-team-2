package repository

import (
	"project_pos_app/model"

	"gorm.io/gorm"
)

type Staff = model.Staff

type StaffRepository interface {
	GetAll() []Staff
	GetByID(id int) *Staff
	Create(staff *Staff)
	Update(staff *Staff)
	Delete(id int)
}

type staffRepository struct {
	DB *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{DB: db}
}

func (repo *staffRepository) GetAll() []Staff {
	var staffs []Staff
	repo.DB.Find(&staffs)
	return staffs
}

func (repo *staffRepository) GetByID(id int) *Staff {
	var staff Staff
	repo.DB.Where("id =?", id).First(&staff)
	return &staff
}

func (repo *staffRepository) Create(staff *Staff) {
	repo.DB.Create(staff)
}

func (repo *staffRepository) Update(staff *Staff) {
	repo.DB.Save(staff)
}

func (repo *staffRepository) Delete(id int) {
	repo.DB.Delete(&Staff{}, id)
}
