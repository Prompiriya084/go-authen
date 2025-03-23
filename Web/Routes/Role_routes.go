package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	repositories "github.com/Prompiriya084/go-authen/Internal/Adapters/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Impl"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RoleSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := repositories.NewRoleRepository(db)
	service := services.NewRoleService(repo)

	handler := handlers.NewRoleHandler(service)
	app.Get("/roles", handler.GetRoleAll)
	app.Post("/role", handler.CreateRole)
}
