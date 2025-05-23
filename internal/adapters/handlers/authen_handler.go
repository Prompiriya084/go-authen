package handlers

import (
	"time"

	request "github.com/Prompiriya084/go-authen/Internal/Adapters/Request"
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Utilities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/gofiber/fiber/v3"
)

type AuthenHandler struct {
	service   services.IAuthService
	validator ports.Validator
}

func NewAuthHandler(service *services.IAuthService, validator ports.Validator) *AuthenHandler {
	return &AuthenHandler{
		service:   *service,
		validator: validator,
	}
}

// UserAuthHandler godoc
// @Summary Hello example
// @Tags UserAuth
// @Accept json
// @Produce json
// @Param userauth body entities.UserAuth true "UserAuth data"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *AuthenHandler) SignIn(c fiber.Ctx) error {
	var request entities.UserAuth
	if err := c.Bind().JSON(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.validator.ValidateStruct(request); err != nil {
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

// UserAuthHandler godoc
// @Summary Hello example
// @Tags UserAuth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param RequestRegister body request.RequestRegister true "RequestRegister data"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /register [post]
func (h *AuthenHandler) Register(c fiber.Ctx) error {
	var request request.RequestRegister
	if err := c.Bind().JSON(&request); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.validator.ValidateStruct(request); err != nil {
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

// UserAuthHandler godoc
// @Summary Hello example
// @Tags UserAuth
// @Accept json
// @Produce json
// @Security CookieAuth
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /signout [post]
func (h *AuthenHandler) SignOut(c fiber.Ctx) error {
	c.ClearCookie("jwt")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully sign out.",
	})
}
