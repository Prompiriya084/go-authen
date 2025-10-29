package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	"github.com/gofiber/fiber/v3"
)

func RoleSetupRouter(app *fiber.App, roleMiddleware *middleware.RoleMiddleware, jwtMiddleware *middleware.JwtMiddleware, roleHandler *handlers.RoleHandler) {
	appRole := app.Group("api/role")
	appRole.Use(jwtMiddleware.AuthMiddleware())
	appRole.Use(roleMiddleware.RequiredRole("admin"))

	appRole.Get("", roleHandler.GetRoleAll)
	appRole.Get("/:id", roleHandler.GetRoleById)
	appRole.Post("", roleHandler.CreateRole)
	appRole.Put("/:id", roleHandler.UpdateRole)
	appRole.Delete("/:id", roleHandler.DeleteRole)
}
