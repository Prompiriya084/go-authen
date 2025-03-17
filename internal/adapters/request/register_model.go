package request

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type RequestRegister struct {
	UserID          uint          `json:"user_id"`
	User            entities.User `json:"user"`
	Role            uint          `json: "role" validate:"required"`
	ConfirmPassword string        `json:"confirm_password" validate:"required"` //validate:"required,confirm_password"
}
