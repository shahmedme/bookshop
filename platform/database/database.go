package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"shakilahmed.com/bookshop/app/models"
)

var DB *gorm.DB

func InitDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	fmt.Println(dbHost, dbPort)

	var err error
	var config string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect database.\n", err)
		os.Exit(2)
	}

	log.Println("Database Connected")
	// DB.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	DB.AutoMigrate(&models.Book{})
}
