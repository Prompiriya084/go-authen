package main

import (
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	web "github.com/Prompiriya084/go-authen/Web/Routes"
	"github.com/Prompiriya084/go-authen/config"
	"github.com/Prompiriya084/go-authen/internal/infrastructure/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	config.LoadEnv()
	db := database.InitDb()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, //Adjust this to be more restrictive if needed such as http://localhost
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	loggingMiddleware := middleware.NewLoggingMiddleware()
	app.Use(loggingMiddleware.Console)

	web.AuthSetupRouter(db, app)
	web.UserSetupRouter(db, app)
	web.RoleSetupRouter(db, app)
	app.Listen(":8080")
}
