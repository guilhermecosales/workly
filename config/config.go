package config

import (
	"errors"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// Initialize the database connection
	return errors.New("not implemented")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
