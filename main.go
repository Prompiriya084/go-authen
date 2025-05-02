package main

import (
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
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

	web.AuthSetupRouter(db, app)
	web.UserSetupRouter(db, app)
	web.RoleSetupRouter(db, app)
	app.Listen(":8080")
}
