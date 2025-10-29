package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"
	"github.com/gofiber/fiber/v3"
)

func UserSetupRouter(app *fiber.App, jwtMiddleware *middleware.JwtMiddleware, userHandler *handlers.UserHandler) {

	appUser := app.Group("api/users")
	appUser.Use(jwtMiddleware.AuthMiddleware())
	appUser.Get("", userHandler.GetUsers)
	appUser.Get("/:id", userHandler.GetUserById)
	appUser.Get("/getByEmail/:email", userHandler.GetUserByEmail)
}
