package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Dodge-git/Test_For_Work/internal/dto"
	"github.com/Dodge-git/Test_For_Work/internal/service"
)

type DepartmentHandler struct {
	service service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service:service}
}
func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req dto.CreateDepartment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	resp, err := h.service.Create(req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, resp)
}
func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := extractID(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	depth := 1
	if d := r.URL.Query().Get("depth"); d != "" {
		if parsed, err := strconv.Atoi(d); err == nil {
			depth = parsed
		}
	}

	includeEmployees := true
	if ie := r.URL.Query().Get("include_employees"); ie != "" {
		includeEmployees = ie == "true"
	}

	resp, err := h.service.Get(id, depth, includeEmployees)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
func (h *DepartmentHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := extractID(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UpdateDepartment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	resp, err := h.service.Update(id, req)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
func (h *DepartmentHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := extractID(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	mode := r.URL.Query().Get("mode")
	reassignStr := r.URL.Query().Get("reassign_to_department_id")

	var reassignTo *uint
	if reassignStr != "" {
		parsed, err := strconv.ParseUint(reassignStr, 10, 64)
		if err == nil {
			val := uint(parsed)
			reassignTo = &val
		}
	}

	if err := h.service.Delete(id, mode, reassignTo); err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}