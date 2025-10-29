package main

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	repositories "github.com/Prompiriya084/go-authen/Internal/Adapters/Repositories"
	utilities "github.com/Prompiriya084/go-authen/Internal/Adapters/Utilities"
	services "github.com/Prompiriya084/go-authen/Internal/Core/Services"
	security "github.com/Prompiriya084/go-authen/Internal/Infrastructure/Security"
	web "github.com/Prompiriya084/go-authen/Web/Routes"
	"github.com/Prompiriya084/go-authen/config"
	_ "github.com/Prompiriya084/go-authen/docs"
	"github.com/Prompiriya084/go-authen/internal/infrastructure/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Fiber v3 Swagger Example
// @version 1.0
// @description Swagger with Fiber v3
// @host localhost:3000
// @BasePath /api

// @securityDefinitions.cookieAuth CookieAuth
// @in cookie
// @name session_token
func main() {

	config.LoadEnv()
	db := database.InitDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, //Adjust this to be more restrictive if needed such as http://localhost
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	// app.Get("/swagger/*", swagger.FiberWrapHandler(httpSwagger.Handler()))
	app.Get("/swagger/*", adaptor.HTTPHandlerFunc(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	)))
	loggingMiddleware := middleware.NewLoggingMiddleware()
	app.Use(loggingMiddleware.Console)

	jwtService := security.NewJwtService()
	jwtMiddleware := middleware.NewJwtMiddleware(jwtService)

	validator := utilities.NewValidator()

	//Repositories
	roleRepo := repositories.NewRoleRepository(db)
	userRepo := repositories.NewUserRepository(db)
	userroleRepo := repositories.NewUserRoleRepository(db)
	userauthRepo := repositories.NewUserAuthRepository(db)
	//Services
	userService := services.NewUserService(userRepo)
	userroleService := services.NewUserRoleService(userroleRepo)
	authService := services.NewAuthService(userRepo, userauthRepo, userroleRepo, roleRepo, jwtService)
	roleService := services.NewRoleService(roleRepo)
	//Handlers
	userHandler := handlers.NewUserHandler(&userService, &validator)
	authHandler := handlers.NewAuthHandler(&authService, &validator)
	roleHandler := handlers.NewRoleHandler(&roleService, &validator)

	roleMiddleware := middleware.NewRoleMiddleware(userroleService)
	//Web routes
	web.UserSetupRouter(app, jwtMiddleware, userHandler)
	web.AuthSetupRouter(app, roleMiddleware, jwtMiddleware, authHandler)
	web.RoleSetupRouter(app, roleMiddleware, jwtMiddleware, roleHandler)
	app.Listen(":8080")
}
