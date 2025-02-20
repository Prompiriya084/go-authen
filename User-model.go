package main

import (
	"gorm.io/gorm"
)

type UserAuth struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Name       string
	Surname    string
	Role       string
	UserAuthID uint //meaning fk
	UserAuth   UserAuth
}
