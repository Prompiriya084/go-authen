package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // Keeps CreatedAt, UpdatedAt, DeletedAt
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"` // Custom Primary Key
	Name       string    `json:"name" validate:"required"`
	Surname    string    `json:"surname" validate:"required"`
	UserAuthID uuid.UUID `gorm:"type:uuid;not null" json:"userauth_id"` //`json:"userauth_id"` //meaning fk
	UserAuth   UserAuth  `gorm:"foreignKey:UserAuthID" json:"userauth"`
}
