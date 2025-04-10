package config

import (
	"fmt"
	"github.com/guilhermecosales/workly/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitializeDatabaseConnection() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database configuration from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		logger.Error("Database configuration is missing in environment variables")
		return nil, fmt.Errorf("missing database configuration")
	}

	// Build the PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = database.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to migrate schema: %v", err)
		return nil, err
	}

	return database, err
}
