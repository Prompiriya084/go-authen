package entities

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
// @Description Represents the User entity in the system
// @type User
type User struct {
	// gorm.Model           // Keeps CreatedAt, UpdatedAt, DeletedAt
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id" swaggerignore:"true"` // Custom Primary Key
	Name       string    `json:"name" validate:"required"`
	Surname    string    `json:"surname" validate:"required"`
	UserAuthID uuid.UUID `gorm:"type:uuid;not null" json:"userauth_id"` //`json:"userauth_id"` //meaning fk
	UserAuth   UserAuth  `gorm:"foreignKey:UserAuthID" json:"userauth"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeleteAt   time.Time `json:"delete_at"`
}
