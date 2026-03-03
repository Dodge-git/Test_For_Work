package repository

import (
	"errors"

	"github.com/Dodge-git/Test_For_Work/internal/models"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(dep *models.Department) error
	Update(dep *models.Department) error
	Delete(id uint) (bool, error)
	GetByID(id uint) (*models.Department, error)
	ListByParent(parentID uint) ([]models.Department, error)
}

type depoRepo struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &depoRepo{db: db}
}

func (r *depoRepo) Create(dep *models.Department) error {
	return r.db.Create(dep).Error
}

func (r *depoRepo) Update(dep *models.Department) error {
	return r.db.Model(&models.Department{}).Where("id = ?", dep.ID).Updates(dep).Error
}

func (r *depoRepo) Delete(id uint) (bool, error) {
	result := r.db.Delete(&models.Department{}, id)

	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (r *depoRepo)ListByParent(parentID uint)([]models.Department,error){
	var departments []models.Department

	err := r.db.Where("parent_id = ?",parentID).Order("created_at ASC").Find(&departments).Error
	return departments,err
}

func (r *depoRepo) GetByID(id uint)(*models.Department,error){
	var department models.Department

	if err := r.db.First(&department,id).Error; err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}
	return &department,nil
}
