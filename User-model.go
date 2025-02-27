package main

import (
	"gorm.io/gorm"
)

//	type RequestRegister struct {
//		UserID          uint
//		User            User
//		UserAuthID      uint
//		UserAuth        UserAuth
//		ConfirmPassword string `json:"confirm_password"`
//	}
type RequestRegister struct {
	UserID uint `json:"user_id"`
	User   User `json:"user"`
	// UserAuthID      uint     `json:"userauth_id"`
	// UserAuth        UserAuth `json:"userauth"`
	ConfirmPassword string `json:"confirm_password" validate:"required,confirm_password"`
}

type UserAuth struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type User struct {
	gorm.Model
	Name       string   `json:"name" validate:"required,name"`
	Surname    string   `json:"surname" validate:"required,surname"`
	Role       string   `json:"role" validation:"required,role"`
	UserAuthID uint     `json:"userauth_id"` //meaning fk
	UserAuth   UserAuth `json:"userauth"`
}
