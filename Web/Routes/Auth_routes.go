package web

import (
	handlers "github.com/Prompiriya084/go-authen/Internal/Adapters/Handlers"
	middleware "github.com/Prompiriya084/go-authen/Internal/Adapters/Middleware"

	"github.com/gofiber/fiber/v3"
)

func AuthSetupRouter(app *fiber.App, roleMiddleware *middleware.RoleMiddleware, jwtMiddleware *middleware.JwtMiddleware, authHandler *handlers.AuthenHandler) {
	appApi := app.Group("api")
	appApi.Post("/login", authHandler.SignIn)
	appApi.Use([]string{"/register", "/signout"}, jwtMiddleware.AuthMiddleware())
	appApi.Use("/register", roleMiddleware.RequiredRole("admin"))
	appApi.Post("/register", authHandler.Register)
	appApi.Post("/signout", authHandler.SignOut)
}
