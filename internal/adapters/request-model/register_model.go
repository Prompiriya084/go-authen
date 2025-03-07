package request

import (
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
)

type RequestRegister struct {
	UserID          uint          `json:"user_id"`
	User            entities.User `json:"user"`
	ConfirmPassword string        `json:"confirm_password" validate:"required"` //validate:"required,confirm_password"
}
