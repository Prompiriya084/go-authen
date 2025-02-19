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
	ID int `gorm:"primaryKey"`
	// UserAuthID int `gorm:"FK"`
	Name    string
	Surname string
	Role    string
}
