package config

import (
	"fmt"
	"log"
	"loyalty-program-api/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// SetupDatabase initializes the database connection and runs AutoMigrate
func InitDB() {
	var err error
	// Read the database URL from the environment variable
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Automigrate the database
	err = DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Redemption{})
	if err != nil {
		log.Fatalf("Error automigrating database: %v", err)
	}

	log.Println("Database connected and tables migrated successfully.")
}
