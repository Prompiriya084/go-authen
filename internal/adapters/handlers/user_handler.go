package handlers

import (
	"fmt"

	request "github.com/Prompiriya084/go-authen/internal/adapters/request-model"
	entities "github.com/Prompiriya084/go-authen/internal/core/entities"
	services "github.com/Prompiriya084/go-authen/internal/core/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(service *services.IUserService) *UserHandler {
	return &UserHandler{service: *service}
}
func (h *UserHandler) Create(c fiber.Ctx) error {
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
	if user, _ := h.service.GetWithUserAuthByEmail(user.UserAuth.Email); user != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data duplicated.",
		})
	}
	user.Role = "user"
	//user.Role = "user"
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(request.User.UserAuth.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user.UserAuth.Password = string(hashedpassword)

	return c.JSON(fiber.Map{
		"message": "register successfully.",
	})
}
func (h *UserHandler) GetUsers(c fiber.Ctx) error {
	users, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(users)
}
