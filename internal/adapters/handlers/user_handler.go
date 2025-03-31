package handlers

import (
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(service *services.IUserService) *UserHandler {
	return &UserHandler{service: *service}
}
func (h *UserHandler) GetUsers(c fiber.Ctx) error {
	users, err := h.service.GetUserAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(users)
}
