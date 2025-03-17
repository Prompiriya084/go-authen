package main

import (
	web "github.com/Prompiriya084/go-authen/Web/Routes"
	"github.com/Prompiriya084/go-authen/config"
	"github.com/Prompiriya084/go-authen/internal/infrastructure/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

func main() {

	config.LoadEnv()
	db := database.InitDb()
	app := fiber.New()

	// app.Post("/register", func(c fiber.Ctx) error {
	// 	var reqRegister entities
	// 	if err := c.Bind().JSON(&reqRegister); err != nil {
	// 		fmt.Print(err.Error())
	// 		return c.SendStatus(fiber.StatusBadRequest)
	// 	}
	// 	//fmt.Println(reqRegister)
	// 	if err := validate.Struct(reqRegister); err != nil {
	// 		return c.SendStatus(fiber.StatusBadRequest)
	// 	}

	// 	var user entities.User
	// 	user = reqRegister.User
	// 	fmt.Println(user)
	// 	if user, _ := getUserWithUserAuthByEmail(db, user.UserAuth.Email); user != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"message": "Data duplicated.",
	// 		})
	// 	}
	// 	user.Role = "user"
	// 	//user.Role = "user"
	// 	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(reqRegister.User.UserAuth.Password), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"message": err.Error(),
	// 		})
	// 	}
	// 	user.UserAuth.Password = string(hashedpassword)

	// 	if err := createUser(db, &user); err != nil {
	// 		fmt.Println(err.Error())
	// 		return c.SendStatus(fiber.StatusInternalServerError)
	// 	}

	// 	return c.JSON(fiber.Map{
	// 		"message": "register successful.",
	// 	})
	// })

	web.AuthSetupRouter(db, app)
	web.UserSetupRouter(db, app)
	web.RoleSetupRouter(db, app)
	app.Listen(":8080")
}
