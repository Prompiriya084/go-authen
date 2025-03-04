package entities

import (
	"gorm.io/gorm"
)

type RequestRegister struct {
	UserID          uint   `json:"user_id"`
	User            User   `json:"user"`
	ConfirmPassword string `json:"confirm_password" validate:"required"` //validate:"required,confirm_password"
}

type User struct {
	gorm.Model
	Name       string   `json:"name" validate:"required"`
	Surname    string   `json:"surname" validate:"required"`
	Role       string   `json:"role"`
	UserAuthID uint     `json:"userauth_id"` //meaning fk
	UserAuth   UserAuth `json:"userauth"`
}
