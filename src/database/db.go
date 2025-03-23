package db

import (
	"fmt"
	"go-cc/src/config"
	"go-cc/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	db, err := gorm.Open(sqlite.Open(config.DatabaseURL()), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect ot database!", err)
	}

	err = db.AutoMigrate(&models.Todo{})

	if err != nil {
		log.Fatal("Failed to migrate ot database!", err)
	}

	DB = db
	fmt.Println("Database connection established and migrated.")
}
