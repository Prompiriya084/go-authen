package handlers

import (
	"fmt"

	ports_utilities "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Utilities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	service   services.IUserService
	validator ports_utilities.Validator
}

func NewUserHandler(service *services.IUserService, validator *ports_utilities.Validator) *UserHandler {
	return &UserHandler{
		service:   *service,
		validator: *validator,
	}
}

// HelloHandler godoc
// @Summary Hello example
// @Tags Users
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 200 {object} entities.User
// @Router /users [get]
func (h *UserHandler) GetUsers(c fiber.Ctx) error {

	users, err := h.service.GetUserAll(nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(users)
}

// RolesHandler godoc
// @Summary Hello example
// @Tags Users
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param   id   path     string  true  "The ID of the resource"
// @Success 200 {object} entities.User
// @Failure 401 {object} entities.User
// @Router /users/{id} [get]
func (h *UserHandler) GetUserById(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// RolesHandler godoc
// @Summary Hello example
// @Tags Users
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param   email   path     string  true  "The email of the resource"
// @Success 200 {object} entities.User
// @Failure 401 {object} entities.User
// @Router /users/getByEmail/{email} [get]
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
