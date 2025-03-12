package web

import (
	"github.com/Prompiriya084/go-authen/internal/adapters/handlers"
	middleware "github.com/Prompiriya084/go-authen/internal/adapters/middleware"
	"github.com/Prompiriya084/go-authen/internal/adapters/repositories"
	services "github.com/Prompiriya084/go-authen/internal/core/services/impl"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func AuthSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := repositories.NewUserRepository(db)
	service := services.NewAuthService(repo)
	authHandler := handlers.NewAuthHandler(&service)

	app.Post("/login", authHandler.SignIn)
	app.Use("/register", middleware.JwtMiddleware("admin"))
	app.Post("/register", authHandler.Register)
}
