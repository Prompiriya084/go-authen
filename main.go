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
const jwtSecretKey = "TestJwtSecretKey" //Should be env

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

	// app.Post("/register", )

	app.Get("/login", func(c fiber.Ctx) error {
		var user UserAuth
		if err := c.Bind().JSON(&user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		token, err := login(db, &user)
		if err != nil {
			fmt.Printf(err.Error())
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.JSON(fiber.Map{
			"message": "login succesful.",
			"token":   token,
		})

	})
}
func login(db *gorm.DB, user *UserAuth) (string, error) {
	var selectedUser UserAuth
	result := db.Where("Email=?", user.Email).First(selectedUser)
	if result.Error != nil {
		fmt.Printf(result.Error.Error())
		return "", result.Error
	}
	hashedpassword := selectedUser.Password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedpassword),
		[]byte(user.Password),
	); err != nil {
		return "", err
	}
	// if err != nil {
	// 	return "", err
	// }

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
