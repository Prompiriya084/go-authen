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

func UserSetupRouter(db *gorm.DB, app *fiber.App) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	jwtService := security.NewJwtService()
	jwtMiddleware := middleware.NewJwtMiddleware(jwtService)

	userHandler := handlers.NewUserHandler(&service)

	appUser := app.Group("api/users")
	appUser.Use(jwtMiddleware.AuthMiddleware())
	appUser.Get("", userHandler.GetUsers)
	appUser.Get("/:id", userHandler.GetUserById)
	appUser.Get("/getByEmail/:email", userHandler.GetUserByEmail)
}
