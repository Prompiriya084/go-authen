package main

import (
	"fmt"

	web "github.com/Prompiriya084/go-authen/Web/Routes"
	"github.com/Prompiriya084/go-authen/config"
	"github.com/Prompiriya084/go-authen/internal/adapters/middleware"
	"github.com/Prompiriya084/go-authen/internal/infrastructure/database"
	"github.com/gofiber/fiber/v3"

	//jwtware "github.com/gofiber/jwt/v3"

	//"gorm.io/driver/postgres"

	//"gorm.io/gorm/logger"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func main() {

	config.LoadEnv()
	db := database.InitDb()

	fmt.Printf("Connect successful.")
	// fmt.Print(db)

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

	// app.Post("/login", func(c fiber.Ctx) error {
	// 	var userAuth entities.UserAuth
	// 	if err := c.Bind().JSON(&userAuth); err != nil {
	// 		return c.SendStatus(fiber.StatusBadRequest)
	// 	}
	// 	fmt.Println(userAuth)
	// 	token, err := login(db, &userAuth)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return c.SendStatus(fiber.StatusUnauthorized)
	// 	}

	// 	return c.JSON(fiber.Map{
	// 		"message": "login succesful.",
	// 		"token":   token,
	// 	})

	// })
	web.AuthSetupRouter(db, app)
	// JWT Middleware
	app.Use(middleware.JwtMiddleware)
	//app.Use("/users", authrequired)
	// app.Get("/users", func(c fiber.Ctx) error {
	// 	return c.JSON(getUsers(db))
	// })

	web.UserSetupRouter(db, app)
	app.Listen(":8080")
}

// func login(db *gorm.DB, user *entities.UserAuth) (string, error) {
// 	var selectedUser entities.UserAuth
// 	result := db.Where("email=?", user.Email).First(&selectedUser)
// 	if result.Error != nil {
// 		fmt.Println(result.Error.Error())
// 		return "", result.Error
// 	}
// 	hashedpassword := selectedUser.Password
// 	fmt.Println(hashedpassword)
// 	fmt.Println(user.Password)
// 	if err := bcrypt.CompareHashAndPassword(
// 		[]byte(hashedpassword),
// 		[]byte(user.Password),
// 	); err != nil {
// 		fmt.Print(err.Error())
// 		return "", err
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"email":      user.Email,
// 		"expireData": time.Now().Add(time.Hour * 1).Unix(),
// 	})
// 	t, err := token.SignedString([]byte(os.Getenv("Jwt_Secret")))
// 	if err != nil {
// 		return "", err
// 	}

// 	return t, nil
// }

// func createUser(db *gorm.DB, user *entities.User) error {
// 	result := db.Create(&user)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
// func getUser(db *gorm.DB, id int) (*entities.User, error) {
// 	var user entities.User
// 	result := db.First(&user, id)
// 	if result.Error != nil {
// 		//log.Fatalf("Error get book: %v", result.Error)
// 		return nil, result.Error
// 	}

// 	return &user, nil
// }
// func getUserWithUserAuthByEmail(db *gorm.DB, email string) (*entities.User, error) {
// 	fmt.Println(email)
// 	// var userAuth entities.UserAuth
// 	// if err := db.Where("email = ?", email).First(&userAuth).Error; err != nil {
// 	// 	fmt.Println("UserAuth not found")
// 	// 	return nil, err
// 	// }

// 	// var user entities.User
// 	// if err := db.Where("user_auth_id = ?", userAuth.ID).Preload("UserAuth").First(&user).Error; err != nil {
// 	// 	fmt.Println("User not found")
// 	// 	return nil, err
// 	// }
// 	var user entities.User
// 	result := db.Preload("UserAuth").Where("user_auths.email = ?", email).Joins("JOIN user_auths ON user_auths.id = users.user_auth_id").First(&user)
// 	if result.Error != nil {
// 		//log.Fatalf("Error get book: %v", result.Error)
// 		return nil, result.Error
// 	}
// 	fmt.Println(user)

// 	return &user, nil
// }
