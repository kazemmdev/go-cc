package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found, using system environment variables.")
	}
}

func AppPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func DatabaseURL() string {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return "./storage/todos.db"
	}
	return dbURL
}
