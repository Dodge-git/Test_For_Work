package repository

import (
	"github.com/Dodge-git/Test_For_Work/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(emp *models.Employee)error
	ListByDepartamentID(depID uint)([]models.Employee,error)
	ReassignDepartment(oldDepID uint,newDepID uint)error
	DeleteByDepartamentID(depID uint)error
}

type empRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &empRepo{db: db}
}

func (r *empRepo) Create(emp *models.Employee) error {
	return r.db.Create(emp).Error
}

func (r *empRepo) ListByDepartamentID(depID uint)([]models.Employee,error){
	var employees []models.Employee

	err := r.db.Where("department_id = ?",depID).Order("created_at ASC").Find(&employees).Error
    return employees,err
}

func (r *empRepo) ReassignDepartment(oldDepID uint,newDepID uint)error{
	return r.db.Model(&models.Employee{}).Where("department_id = ?",oldDepID).Update("department_id",newDepID).Error
 }

 func (r *empRepo) DeleteByDepartamentID(depID uint)error{
	return r.db.Where("department_id = ?",depID).Delete(&models.Employee{}).Error
 }