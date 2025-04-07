package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAuth struct {
	// ID        string         `gorm:"primary_key;" json:"id"` // Custom Primary Key
	// CreatedAt time.Time      // Inherited from gorm.Model
	// UpdatedAt time.Time      // Inherited from gorm.Model
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft Delete
	gorm.Model           // Keeps CreatedAt, UpdatedAt, DeletedAt
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`              // Custom Primary Key
	Email      string    `json:"email" gorm:"unique" validate:"required,email"` //must send with xxx@xxxx
	Password   string    `json:"password" validate:"required"`
}
