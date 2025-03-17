package entities

type UserRole struct {
	UserID uint `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`
	RoleID uint `gorm:"column:role_id"`
	Role   Role `gorm:"foreignKey:RoleID"`
}
