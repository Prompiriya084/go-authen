package web

import (
	handler "github.com/Prompiriya084/go-authen/internal/adapters/handlers"
	adapters "github.com/Prompiriya084/go-authen/internal/adapters/repositories"
	core "github.com/Prompiriya084/go-authen/internal/core/services"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := adapters.NewUserRepository(db)
	service := core.NewUserService(repo)
	userHandler := handler.NewUserHandler(&service)
	app.Get("/users", userHandler.GetUsers)
}
