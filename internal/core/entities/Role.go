package entities

import "gorm.io/gorm"

type Role struct {
	// ID   string `json:"role_id" validate:"required"`
	gorm.Model
	Name string `json:"name" gorm:"unique" validate:"required"`
}
