package handlers

import (
	"github.com/Prompiriya084/go-authen/internal/core/entities"
	"github.com/gofiber/fiber/v3"

	request "github.com/Prompiriya084/go-authen/internal/adapters/request"
	services "github.com/Prompiriya084/go-authen/internal/core/services/interfaces"
	"github.com/go-playground/validator/v10"
)

type AuthenHandler struct {
	service services.IAuthService
}

func NewAuthHandler(service *services.IAuthService) *AuthenHandler {
	return &AuthenHandler{service: *service}
}
func (h *AuthenHandler) SignIn(c fiber.Ctx) error {
	var request entities.UserAuth
	if err := c.Bind().JSON(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var validate = validator.New()
	if err := validate.Struct(request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := h.service.SignIn(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "login successful.",
		"token":   token,
	})
}
func (h *AuthenHandler) Register(c fiber.Ctx) error {
	var request request.RequestRegister

	if err := c.Bind().JSON(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var validate = validator.New()
	if err := validate.Struct(request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var user entities.User
	user = request.User

	if err := h.service.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "register successfully.",
	})
}
