package handlers

import (
	"time"

	request "github.com/Prompiriya084/go-authen/Internal/Adapters/Request"
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Interfaces"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
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

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1), // for cookies only not jwt
		HTTPOnly: true,                          // localhost
	})

	return c.JSON(fiber.Map{
		"message": "login successful.",
		//"token":   token,
	})
}
func (h *AuthenHandler) Register(c fiber.Ctx) error {
	var request request.RequestRegister
	// localsUserId := c.Locals("user_id")
	// userId, _ := localsUserId.(uint)
	// fmt.Println(userId)
	if err := c.Bind().JSON(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var validate = validator.New()
	if err := validate.Struct(request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if request.Password != request.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).SendString("Password and confirm password is not equal.")
	}
	if err := h.service.Register(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "register successfully.",
	})
}
