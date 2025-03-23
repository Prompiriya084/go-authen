package entities

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `gorm:"column:user_id"`
	User   User      `gorm:"foreignKey:UserID"`
	RoleID uint      `gorm:"column:role_id"`
	Role   Role      `gorm:"foreignKey:RoleID"`
}
