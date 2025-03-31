package entities

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID   uint   `gorm:"type:uint;primary_key;" json:"id"`
	Name string `json:"name" gorm:"unique" validate:"required"`
}
