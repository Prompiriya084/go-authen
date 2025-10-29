package handlers

import (
	"strconv"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports_utilities "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Utilities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	"github.com/gofiber/fiber/v3"
)

type RoleHandler struct {
	service   services.IRoleService
	validator ports_utilities.Validator
}

func NewRoleHandler(service *services.IRoleService, validator *ports_utilities.Validator) *RoleHandler {
	return &RoleHandler{
		service:   *service,
		validator: *validator,
	}
}

// RolesHandler godoc
// @Summary Hello example
// @Tags Role
// @Accept  json
// @Produce json
// @Security CookieAuth
// @Success 200 {object} entities.Role
// @Failure 401 {object} entities.Role
// @Router /role [get]
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

// RolesHandler godoc
// @Summary Hello example
// @Tags Role
// @Accept json
// @Produce json
// @Param   id   path     string  true  "The ID of the resource"
// @Security CookieAuth
// @Success 200 {object} entities.Role
// @Failure 401 {object} entities.Role
// @Router /role/{id} [get]
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

// CreateUser godoc
// @Summary Hello example
// @Tags Role
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param role body entities.Role true "Role data"
// @Success 200 {object} entities.Role
// @Failure 401 {object} entities.Role
// @Router /role [post]
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

// RolesHandler godoc
// @Summary Hello example
// @Tags Role
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param   id   path     string  true  "The ID of the resource"
// @Param role body entities.Role true "Role data"
// @Success 200 {object} entities.Role
// @Failure 401 {object} entities.Role
// @Router /role/{id} [put]
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

// RolesHandler godoc
// @Summary Hello example
// @Tags Role
// @Accept json
// @Produce json
// @Security CookieAuth
// @Param   id   path     string  true  "The ID of the resource"
// @Success 200 {object} entities.Role
// @Failure 401 {object} entities.Role
// @Router /role/{id} [delete]
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
