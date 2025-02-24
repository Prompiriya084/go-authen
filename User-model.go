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
	ConfirmPassword string `json:"confirm_password"`
}

type UserAuth struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Name       string   `json:"name"`
	Surname    string   `json:"surname"`
	Role       string   `json:"role"`
	UserAuthID uint     `json:"userauth_id"` //meaning fk
	UserAuth   UserAuth `json:"userauth"`
}
