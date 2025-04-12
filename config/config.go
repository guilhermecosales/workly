package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize the database connection
	db, err = InitializeDatabaseConnection()

	if err != nil {
		return fmt.Errorf("error initializing database connection: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetDatabase() *gorm.DB {
	return db
}
