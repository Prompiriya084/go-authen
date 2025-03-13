package web

import (
	handler "github.com/Prompiriya084/go-authen/internal/adapters/handlers"
	middleware "github.com/Prompiriya084/go-authen/internal/adapters/middleware"
	implRepositories "github.com/Prompiriya084/go-authen/internal/adapters/repositories"
	services "github.com/Prompiriya084/go-authen/internal/core/services/impl"
	"github.com/Prompiriya084/go-authen/internal/infrastructure/security"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := implRepositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	userHandler := handler.NewUserHandler(&service)

	jwtService := security.NewJwtService()
	jwtMiddleware := middleware.JwtMiddleware(&jwtService, repo)
	// JWT Middleware
	app.Use("/users", jwtMiddleware)
	app.Get("/users", userHandler.GetUsers)
}
