package service

import (
	"strings"
	"time"

	"github.com/Dodge-git/Test_For_Work/internal/dto"
	"github.com/Dodge-git/Test_For_Work/internal/models"
	"github.com/Dodge-git/Test_For_Work/internal/repository"
)

type EmployeeService interface {
	Create(departmentID uint, req dto.CreateEmployee) (*dto.EmployeeResponse, error)
}

type employeeService struct {
	depRepo repository.DepartmentRepository
	empRepo repository.EmployeeRepository
}

func NewEmployeeService(depRepo repository.DepartmentRepository, empRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		depRepo: depRepo,
		empRepo: empRepo,
	}
}

func (s *employeeService) Create(departmentID uint, req dto.CreateEmployee) (*dto.EmployeeResponse, error) {

	dep, err := s.depRepo.GetByID(departmentID)
	if err != nil {
		return nil, err
	}
	if dep == nil {
		return nil, ErrDepartmentNotFound
	}

	fullName := strings.TrimSpace(req.FullName)
	position := strings.TrimSpace(req.Position)

	if len(fullName) < 1 || len(fullName) > 200 {
		return nil, ErrInvalidFullName
	}

	if len(position) < 1 || len(position) > 200 {
		return nil, ErrInvalidPosition
	}

	model := &models.Employee{
		DepartmentID: departmentID,
		FullName:     fullName,
		Position:     position,
		HiredAt:      req.HiredAt,
		CreatedAt:    time.Now(),
	}

	if err := s.empRepo.Create(model); err != nil {
		return nil, err
	}

	resp := &dto.EmployeeResponse{
		ID:           model.ID,
		DepartmentID: model.DepartmentID,
		FullName:     model.FullName,
		Position:     model.Position,
		HiredAt:      model.HiredAt,
		CreatedAt:    model.CreatedAt,
	}

	return resp, nil
}