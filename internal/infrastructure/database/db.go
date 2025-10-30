package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() *gorm.DB {

	host := os.Getenv("DB_Host")
	port, _ := strconv.Atoi(os.Getenv("DB_Port"))
	username := os.Getenv("DB_Username")
	password := os.Getenv("DB_Password")
	database := os.Getenv("DB_Name")

	// dsn := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, username, password, database)

	// Connection string (same format as sqlserver driver)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable",
		username, password, host, port, database)

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
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: newLogger,
	// })

	// Connect using GORM
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✅ Connected to SQL Server")

	if err != nil {
		panic("Failed to connect database.")
	}
	fmt.Printf("Connect successful.")

	db.AutoMigrate(entities.User{}, entities.UserAuth{}, entities.UserRole{}, entities.Role{})

	return db
}
