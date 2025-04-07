package handlers

import (
	"fmt"
	"strconv"

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
func (h *UserHandler) GetUserById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user, err := h.service.GetUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}
func (h *UserHandler) GetUserByEmail(c fiber.Ctx) error {
	fmt.Println("user_handler Get email method!!!!")
	email := c.Params("email")
	fmt.Println(email)
	if email == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Email is required.")
	}
	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"name":    user.Name,
			"surname": user.Surname,
		},
	})
}
