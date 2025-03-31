package main

import (
	web "github.com/Prompiriya084/go-authen/Web/Routes"
	"github.com/Prompiriya084/go-authen/config"
	"github.com/Prompiriya084/go-authen/internal/infrastructure/database"
	"github.com/gofiber/fiber/v3"
)

func main() {

	config.LoadEnv()
	db := database.InitDb()
	app := fiber.New()

	web.AuthSetupRouter(db, app)
	web.UserSetupRouter(db, app)
	web.RoleSetupRouter(db, app)
	app.Listen(":8080")
}
