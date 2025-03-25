package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	repositories "github.com/Prompiriya084/go-authen/Internal/Adapters/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Impl"
	security "github.com/Prompiriya084/go-authen/Internal/Infrastructure/Security"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RoleSetupRouter(db *gorm.DB, app *fiber.App) {
	roleRepo := repositories.NewRoleRepository(db)
	userRoleRepo := repositories.NewUserRoleRepository(db)

	roleService := services.NewRoleService(roleRepo)
	userRoleService := services.NewUserRoleService(userRoleRepo)
	jwtService := security.NewJwtService()
	jwtMiddleware := middleware.NewJwtMiddleware(jwtService)

	roleMiddleware := middleware.NewRoleMiddleware(userRoleService)

	handler := handlers.NewRoleHandler(roleService)
	app.Use("/role", jwtMiddleware.AuthMiddleware())
	app.Get("/role", handler.GetRoleAll)
	app.Post("/role", roleMiddleware.RequiredRole("admin"), handler.CreateRole)
}
