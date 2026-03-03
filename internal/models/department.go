package models

import (
	"time"
)

type Department struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:200;not null"`
	ParentID  *uint     `json:"parent_id" gorm:"index"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:now()"`

	Parent *Department `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}