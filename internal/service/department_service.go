package service

import (
	"strings"

	"github.com/Dodge-git/Test_For_Work/internal/dto"
	"github.com/Dodge-git/Test_For_Work/internal/models"
	"github.com/Dodge-git/Test_For_Work/internal/repository"
)

type DepartmentService interface {
	Create(req dto.CreateDepartment) (*dto.DepartmentResponse, error)
	Get(id uint, depth int, includeEmployees bool) (*dto.DepartmentResponse, error)
	Update(id uint, req dto.UpdateDepartment) (*dto.DepartmentResponse, error)
	Delete(id uint, mode string, reassignTo *uint) error

}

type DepoService struct {
	depoRepo repository.DepartmentRepository
	empRepo  repository.EmployeeRepository
	tx       repository.TransactionManager
}

func NewDepartmentService(depoRepo repository.DepartmentRepository, empRepo repository.EmployeeRepository, tx repository.TransactionManager) DepartmentService {
	return &DepoService{
		depoRepo: depoRepo,
		empRepo:  empRepo,
		tx:       tx,
	}
}

func (s *DepoService) Create(req dto.CreateDepartment) (*dto.DepartmentResponse, error) {

	name := strings.TrimSpace(req.Name)

	if len(name) < 1 || len(name) > 200 {
		return nil, ErrInvalidName
	}

	if req.ParentID != nil {
		parent, err := s.depoRepo.GetByID(*req.ParentID)
		if err != nil {
			return nil, err
		}
		if parent == nil {
			return nil, ErrParentNotFound
		}
	}

	var siblings []models.Department
	var err error

	if req.ParentID == nil {
		siblings, err = s.depoRepo.ListByParent(0)
	} else {
		siblings, err = s.depoRepo.ListByParent(*req.ParentID)
	}

	if err != nil {
		return nil, err
	}

	for _, d := range siblings {
		if strings.EqualFold(strings.TrimSpace(d.Name), name) {
			return nil, ErrDuplicateName
		}
	}

	model := &models.Department{
		Name:     name,
		ParentID: req.ParentID,
	}

	if err := s.depoRepo.Create(model); err != nil {
		return nil, err
	}

	resp := &dto.DepartmentResponse{
		ID:        model.ID,
		Name:      model.Name,
		ParentID:  model.ParentID,
		CreatedAt: model.CreatedAt,
	}

	return resp, nil
}

func (s *DepoService) Get(id uint, depth int, includeEmployees bool)(*dto.DepartmentResponse, error){
	if depth < 1 {
		depth = 1
	}

	if depth > 5 {
		return nil, ErrDepthExceeded
	}

	dep, err := s.depoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dep == nil {
		return nil, ErrDepartmentNotFound
	}

	return s.buildTree(dep, depth, includeEmployees)
}
func (s *DepoService) buildTree(dep *models.Department, depth int, includeEmployees bool,) (*dto.DepartmentResponse, error) {
	resp := &dto.DepartmentResponse{
		ID:        dep.ID,
		Name:      dep.Name,
		ParentID:  dep.ParentID,
		CreatedAt: dep.CreatedAt,
	}

	if includeEmployees {
		employees, err := s.empRepo.ListByDepartamentID(dep.ID)
		if err != nil {
			return nil,err
		}
		
		for _, r := range employees {
			resp.Employees = append(resp.Employees, dto.EmployeeResponse{
				ID:           r.ID,
				DepartmentID: r.DepartmentID,
				FullName:     r.FullName,
				Position:     r.Position,
				HiredAt:      r.HiredAt,
				CreatedAt:    r.CreatedAt,
			})
		}
	}
   
	if depth == 1 {
		return resp, nil
	}

	children, err := s.depoRepo.ListByParent(dep.ID) 
	if err != nil {
		return nil,err
	}
	for _, child := range children {
		childCopy := child
		childResp, err := s.buildTree(&childCopy, depth-1, includeEmployees)
		if err != nil{
			return nil,err
		}
		resp.Children = append(resp.Children, *childResp)
	}

	return resp, nil
}
func (s *DepoService) Update(id uint, req dto.UpdateDepartment) (*dto.DepartmentResponse, error) {

	dep, err := s.depoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if dep == nil {
		return nil, ErrDepartmentNotFound
	}

	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)

		if len(name) < 1 || len(name) > 200 {
			return nil, ErrInvalidName
		}

		var siblings []models.Department
		if dep.ParentID == nil {
			siblings, err = s.depoRepo.ListByParent(0)
		} else {
			siblings, err = s.depoRepo.ListByParent(*dep.ParentID)
		}
		if err != nil {
			return nil, err
		}

		for _, sibl := range siblings {
			if sibl.ID != dep.ID &&
				strings.EqualFold(strings.TrimSpace(sibl.Name), name) {
				return nil, ErrDuplicateName
			}
		}

		dep.Name = name
	}

	if req.ParentID != nil {

		if *req.ParentID == dep.ID {
			return nil, ErrInvalidParent
		}

		if *req.ParentID != 0 {
			parent, err := s.depoRepo.GetByID(*req.ParentID)
			if err != nil {
				return nil, err
			}
			if parent == nil {
				return nil, ErrParentNotFound
			}
		}

		if *req.ParentID != 0 {
			isCycle, err := s.isDescendant(dep.ID, *req.ParentID)
			if err != nil {
				return nil, err
			}
			if isCycle {
				return nil, ErrCycleDetected
			}
		}

		if *req.ParentID == 0 {
			dep.ParentID = nil
		} else {
			dep.ParentID = req.ParentID
		}
	}

	if err := s.depoRepo.Update(dep); err != nil {
		return nil, err
	}

	resp := &dto.DepartmentResponse{
		ID:        dep.ID,
		Name:      dep.Name,
		ParentID:  dep.ParentID,
		CreatedAt: dep.CreatedAt,
	}

	return resp, nil
}
func (s *DepoService)Delete(id uint, mode string, reassignTo *uint) error{

	dep, err := s.depoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if dep == nil {
		return ErrDepartmentNotFound
	}

	switch mode {

	case "cascade":
		_, err := s.depoRepo.Delete(id)
		return err

	case "reassign":

		if reassignTo == nil {
			return ErrReassignRequired
		}

		if *reassignTo == id {
			return ErrInvalidParent
		}

		newDep, err := s.depoRepo.GetByID(*reassignTo)
		if err != nil {
			return err
		}
		if newDep == nil {
			return ErrParentNotFound
		}

		return s.tx.NewTransaction(func() error {
			
			if err := s.empRepo.ReassignDepartment(id, *reassignTo); err != nil {
				return err
			}

			_, err := s.depoRepo.Delete(id)
			return err
		})

	default:
		return ErrInvalidDeleteMode
	}
}

func (s *DepoService) isDescendant(parentID uint, targetID uint) (bool, error) {

	children, err := s.depoRepo.ListByParent(parentID)
	if err != nil {
		return false, err
	}

	for _, child := range children {

		if child.ID == targetID {
			return true, nil
		}

		found, err := s.isDescendant(child.ID, targetID)
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
	}

	return false, nil
}