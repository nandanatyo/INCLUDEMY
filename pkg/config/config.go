package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvirontment() {
	err := godotenv.Load()
	env := os.Getenv(".env")
	if err != nil && env == "" {
		log.Println("Error loading .env file")
	}
}
