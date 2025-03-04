package entities

import (
	"gorm.io/gorm"
)

type UserAuth struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
