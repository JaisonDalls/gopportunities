package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Init is ...
func Init() error {
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Error initializing Sqlite: %v", err)
	}
	return nil
}

// GetSQLite is ...
func GetSQLite() *gorm.DB {
	return db
}

// GetLogger is ...
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
