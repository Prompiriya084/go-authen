package entities

import (
	"time"

	"github.com/google/uuid"
)

// User represents a userauths in the system
// @Description Represents the User entity in the system
// @type UserAuth
type UserAuth struct {
	// gorm.Model           // Keeps CreatedAt, UpdatedAt, DeletedAt
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id" swaggerignore:"true"` // Custom Primary Key
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`         //must send with xxx@xxxx
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
	DeleteAt  time.Time `json:"delete_at" swaggerignore:"true"`
}
