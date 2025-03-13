package web

import (
	"github.com/Prompiriya084/go-authen/internal/adapters/handlers"
	middleware "github.com/Prompiriya084/go-authen/internal/adapters/middleware"
	"github.com/Prompiriya084/go-authen/internal/adapters/repositories"
	services "github.com/Prompiriya084/go-authen/internal/core/services/impl"
	security "github.com/Prompiriya084/go-authen/internal/infrastructure/security"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func AuthSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := repositories.NewUserRepository(db)
	jwtService := security.NewJwtService()
	service := services.NewAuthService(repo, jwtService)
	authHandler := handlers.NewAuthHandler(&service)
	app.Post("/login", authHandler.SignIn)
	//app.Use("/register", )
	app.Post("/register", middleware.JwtMiddleware(&jwtService, repo), middleware.RoleMiddleware("admin"), authHandler.Register)
}
