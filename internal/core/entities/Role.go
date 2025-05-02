package entities

import (
	"time"
)

// Role represents a role in the system
// @Description Represents the Role entity in the system
// @type Role
type Role struct {
	// gorm.Model
	ID        uint      `gorm:"type:uint;primary_key;" json:"id" swaggerignore:"true"`
	Name      string    `json:"name" gorm:"unique" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at"`
}
