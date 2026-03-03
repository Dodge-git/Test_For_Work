package transport

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dodge-git/Test_For_Work/internal/dto"
)

type mockDepartmentService struct{}

func (m *mockDepartmentService) Create(req dto.CreateDepartment) (*dto.DepartmentResponse, error) {
	return &dto.DepartmentResponse{
		ID:   1,
		Name: req.Name,
	}, nil
}
func (m *mockDepartmentService) Get(id uint, depth int, includeEmployees bool) (*dto.DepartmentResponse, error) {
	return nil, nil
}
func (m *mockDepartmentService) Update(id uint, req dto.UpdateDepartment) (*dto.DepartmentResponse, error) {
	return nil, nil
}
func (m *mockDepartmentService) Delete(id uint, mode string, reassignTo *uint) error {
	return nil
}


func TestCreateDepartment(t *testing.T) {

	mockService := &mockDepartmentService{}
	handler := NewDepartmentHandler(mockService)

	body := []byte(`{"name":"Backend"}`)
	req := httptest.NewRequest(http.MethodPost, "/departments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.CreateDepartment(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}