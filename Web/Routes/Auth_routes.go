package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	utilities "github.com/Prompiriya084/go-authen/Internal/Adapters/Utilities"

	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	security "github.com/Prompiriya084/go-authen/Internal/Infrastructure/Security"

	repositories "github.com/Prompiriya084/go-authen/Internal/Adapters/Repositories"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func AuthSetupRouter(db *gorm.DB, app *fiber.App) {
	repoUser := repositories.NewUserRepository(db)
	repoUserAuth := repositories.NewUserAuthRepository(db)
	repoUserRole := repositories.NewUserRoleRepository(db)
	repoRole := repositories.NewRoleRepository(db)
	jwtService := security.NewJwtService()

	validatorService := utilities.NewValidator()

	serviceAuth := services.NewAuthService(repoUser, repoUserAuth, repoUserRole, repoRole, jwtService)
	servicesUserRole := services.NewUserRoleService(repoUserRole)

	jwtMiddleware := middleware.NewJwtMiddleware(jwtService)
	roleMiddleware := middleware.NewRoleMiddleware(servicesUserRole)
	authHandler := handlers.NewAuthHandler(&serviceAuth, validatorService)

	appApi := app.Group("api")
	appApi.Post("/login", authHandler.SignIn)
	appApi.Use([]string{"/register", "/signout"}, jwtMiddleware.AuthMiddleware())
	appApi.Use("/register", roleMiddleware.RequiredRole("admin"))
	appApi.Post("/register", authHandler.Register)
	appApi.Post("/signout", authHandler.SignOut)
}
