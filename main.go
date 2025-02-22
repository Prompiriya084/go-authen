package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"

	//"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	username = "myuser"
	password = "mypassword"
	database = "mydatabase"
)

var jwtSecretKey = []byte("TestJwtSecretKey") //Should be env

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
	db.AutoMigrate(User{}, UserAuth{})

	app := fiber.New()

	app.Post("/register", func(c fiber.Ctx) error {
		var user User
		if err := c.Bind().JSON(&user); err != nil {
			fmt.Print(err.Error())
			return c.SendStatus(fiber.StatusBadRequest)
		}
		//user.Role = "admin"
		user.Role = "user"
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.UserAuth.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		user.UserAuth.Password = string(hashedpassword)

		if err := createUser(db, &user); err != nil {
			fmt.Println(err.Error())
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"message": "register successful.",
		})
	})

	app.Post("/login", func(c fiber.Ctx) error {
		var user UserAuth
		if err := c.Bind().JSON(&user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		token, err := login(db, &user)
		if err != nil {
			fmt.Println(err.Error())
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.JSON(fiber.Map{
			"message": "login succesful.",
			"token":   token,
		})

	})

	app.Listen(":8080")
}
func login(db *gorm.DB, user *UserAuth) (string, error) {
	var selectedUser UserAuth
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
	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func createUser(db *gorm.DB, user *User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func getUsers(db *gorm.DB) []User {
	var user []User
	result := db.Find(&user)
	if result.Error != nil {
		fmt.Println("Error get books: %v", result.Error)
	}
	return user
}
func getUser(db *gorm.DB, id int) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		//log.Fatalf("Error get book: %v", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
