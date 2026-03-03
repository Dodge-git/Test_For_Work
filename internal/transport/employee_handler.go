package transport

import (
	"encoding/json"
	"net/http"

	"github.com/Dodge-git/Test_For_Work/internal/dto"
	"github.com/Dodge-git/Test_For_Work/internal/service"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	departmentID, err := extractDepartmentIDFromEmployeePath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid department id")
		return
	}

	var req dto.CreateEmployee
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	resp, err := h.service.Create(departmentID, req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}