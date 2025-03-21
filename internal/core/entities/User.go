package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string   `json:"name" validate:"required"`
	Surname    string   `json:"surname" validate:"required"`
	UserAuthID uint     `json:"userauth_id"` //meaning fk
	UserAuth   UserAuth `json:"userauth"`
}
