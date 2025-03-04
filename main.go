package main

import (
	"fmt"
	"log"
	"os"
	"time"

	//"String"
	entities "github.com/Prompiriya084/go-authen/Entities"
	"github.com/gofiber/fiber/v3"

	//jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/go-playground/validator/v10"
)

const (
	host     = "localhost"
	port     = 5432
	username = "myuser"
	password = "mypassword"
	database = "mydatabase"
)

// var jwtSecretKey = []byte("TestJwtSecretKey") //Should be env
var validate = validator.New()

func authrequired(c fiber.Ctx) error {
	// authHeader := c.Get("Authorization") // Get Authorization header
	// if authHeader == "" {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }
	// // Ensure it starts with "Bearer "
	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// if tokenString == authHeader { // Token not prefixed with "Bearer "
	// 	return c.SendStatus(fiber.StatusUnauthorized) //.JSON(fiber.Map{"error": "Invalid token format"})
	// }

	// token, err := jwt.ParseWithClaims(tokenString, jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("Jwt_Secret")), nil
	// })
	// if err != nil || !token.Valid {
	// 	fmt.Print(err.Error())
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }
	tokenString := c.Get("Authorization") // Get token from header
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	// Remove "Bearer " prefix if present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Jwt_Secret")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}
	// token, err := jwt.ParseWithClaims(tokenString, jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("Jwt_Secret")), nil
	// })
	// if err != nil || !token.Valid {
	// 	fmt.Print(err.Error())
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }

	// jwtware.New(jwtware.Config{
	// 	SigningKey: []byte(os.Getenv("Jwt_Secret")),
	// 	// ErrorHandler: ,
	// })

	return c.Next()
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			// IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect database.")
	}

	fmt.Printf("Connect successful.")
	// fmt.Print(db)
	db.AutoMigrate(entities.User{}, entities.UserAuth{})

	app := fiber.New()

	app.Post("/register", func(c fiber.Ctx) error {
		var reqRegister entities.RequestRegister
		if err := c.Bind().JSON(&reqRegister); err != nil {
			fmt.Print(err.Error())
			return c.SendStatus(fiber.StatusBadRequest)
		}
		//fmt.Println(reqRegister)
		if err := validate.Struct(reqRegister); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var user entities.User
		user = reqRegister.User
		fmt.Println(user)
		if user, _ := getUserWithUserAuthByEmail(db, user.UserAuth.Email); user != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Data duplicated.",
			})
		}
		user.Role = "user"
		//user.Role = "user"
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(reqRegister.User.UserAuth.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		user.UserAuth.Password = string(hashedpassword)

		if err := createUser(db, &user); err != nil {
			fmt.Println(err.Error())
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"message": "register successful.",
		})
	})

	app.Post("/login", func(c fiber.Ctx) error {
		var userAuth entities.UserAuth
		if err := c.Bind().JSON(&userAuth); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		fmt.Println(userAuth)
		token, err := login(db, &userAuth)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.JSON(fiber.Map{
			"message": "login succesful.",
			"token":   token,
		})

	})
	// JWT Middleware
	app.Use(authrequired)
	//app.Use("/users", authrequired)
	app.Get("/users", func(c fiber.Ctx) error {
		return c.JSON(getUsers(db))
	})
	app.Listen(":8080")
}
func login(db *gorm.DB, user *entities.UserAuth) (string, error) {
	var selectedUser entities.UserAuth
	result := db.Where("email=?", user.Email).First(&selectedUser)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return "", result.Error
	}
	hashedpassword := selectedUser.Password
	fmt.Println(hashedpassword)
	fmt.Println(user.Password)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedpassword),
		[]byte(user.Password),
	); err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"expireData": time.Now().Add(time.Hour * 1).Unix(),
	})
	t, err := token.SignedString([]byte(os.Getenv("Jwt_Secret")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func createUser(db *gorm.DB, user *entities.User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func getUsers(db *gorm.DB) []entities.User {
	var user []entities.User
	result := db.Find(&user)
	if result.Error != nil {
		fmt.Printf("Error get books: %v", result.Error)
	}
	return user
}
func getUser(db *gorm.DB, id int) (*entities.User, error) {
	var user entities.User
	result := db.First(&user, id)
	if result.Error != nil {
		//log.Fatalf("Error get book: %v", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
func getUserWithUserAuthByEmail(db *gorm.DB, email string) (*entities.User, error) {
	fmt.Println(email)
	// var userAuth entities.UserAuth
	// if err := db.Where("email = ?", email).First(&userAuth).Error; err != nil {
	// 	fmt.Println("UserAuth not found")
	// 	return nil, err
	// }

	// var user entities.User
	// if err := db.Where("user_auth_id = ?", userAuth.ID).Preload("UserAuth").First(&user).Error; err != nil {
	// 	fmt.Println("User not found")
	// 	return nil, err
	// }
	var user entities.User
	result := db.Preload("UserAuth").Where("user_auths.email = ?", email).Joins("JOIN user_auths ON user_auths.id = users.user_auth_id").First(&user)
	if result.Error != nil {
		//log.Fatalf("Error get book: %v", result.Error)
		return nil, result.Error
	}
	fmt.Println(user)

	return &user, nil
}
