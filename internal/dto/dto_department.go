package dto

import "time"

type DepartmentResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ParentID  *uint     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`

	Employees []EmployeeResponse   `json:"employees,omitempty"`
	Children  []DepartmentResponse `json:"children,omitempty"`
}

type CreateDepartment struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}

type UpdateDepartment struct {
	Name     *string `json:"name"`
	ParentID *uint   `json:"parent_id"`
}