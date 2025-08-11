package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	
	// Set up logger for GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open("jlgo.db"), &gorm.Config{
		Logger: newLogger,
	})
	
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	log.Println("Database connection established")
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return DB
}
