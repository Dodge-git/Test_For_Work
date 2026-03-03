package models

import (
	"time"
)

type Employee struct {
	ID           uint       `gorm:"primaryKey"`
	DepartmentID uint       `json:"department_id" gorm:"not null;index"`
	FullName     string     `json:"full_name" gorm:"size:200;not null"`
	Position     string     `json:"position" gorm:"size:200;not null"`
	HiredAt      *time.Time `json:"hired_at"`
	CreatedAt    time.Time  `json:"created_at" gorm:"not null;default:now()"`

	Department Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}