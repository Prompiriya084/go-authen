package handlers

import (
	"strconv"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Utilities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type RoleHandler struct {
	service   services.IRoleService
	validator ports.Validator
}

var validate = validator.New()

func NewRoleHandler(service services.IRoleService, validator ports.Validator) *RoleHandler {
	return &RoleHandler{
		service:   service,
		validator: validator,
	}
}

func (h *RoleHandler) GetRoleAll(c fiber.Ctx) error {
	roles, err := h.service.GetRoleAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if roles == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{
		"data": roles,
	})
}
func (h *RoleHandler) GetRoleById(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 0, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userId := uint(id)

	roles, err := h.service.GetRole(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Data not found")
	}
	if roles == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{
		"data": roles,
	})
}
func (h *RoleHandler) CreateRole(c fiber.Ctx) error {
	var role *entities.Role
	if err := c.Bind().JSON(&role); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.validator.ValidateStruct(role); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.service.CreateRole(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "create role successfully.",
	})

}
func (h *RoleHandler) UpdateRole(c fiber.Ctx) error {
	roleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var role *entities.Role
	if err := c.Bind().JSON(&role); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.validator.ValidateStruct(role); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid role data: " + err.Error())
	}
	role.ID = uint(roleId)
	if err := h.service.UpdateRole(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Role updated successfully.",
	})
}
func (h *RoleHandler) DeleteRole(c fiber.Ctx) error {
	roleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := h.service.DeleteRole(uint(roleId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "role deleted successfully.",
	})

}
