package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	repositories "github.com/Prompiriya084/go-authen/Internal/Adapters/Repositories"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services/Impl"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(&service)

	// jwtService := security.NewJwtService()
	// jwtMiddleware := middleware.NewJwtMiddleware(jwtService)
	// JWT Middleware
	//app.Use("/users", jwtMiddleware.AuthMiddleware())
	app.Get("/users", userHandler.GetUsers)
}
