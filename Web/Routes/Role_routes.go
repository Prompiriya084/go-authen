package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	repositories "github.com/Prompiriya084/go-authen/Internal/Adapters/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
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

	appRole := app.Group("/role")
	appRole.Use(jwtMiddleware.AuthMiddleware())
	appRole.Use(roleMiddleware.RequiredRole("admin"))
	appRole.Get("", handler.GetRoleAll)
	appRole.Get("/:id", handler.GetRoleById)
	appRole.Post("", handler.CreateRole)
}
