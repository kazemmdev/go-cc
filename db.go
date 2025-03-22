package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect ot database!", err)
	}

	err = db.AutoMigrate(&Todo{})

	if err != nil {
		log.Fatal("Failed to migrate ot database!", err)
	}

	DB = db
}
