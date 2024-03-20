package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
)

func LoadEnvirontment() {
	err := godotenv.Load()
	env := os.Getenv(".env")
	if err != nil && env == "" {
		log.Println("Error loading .env file")
	}
}

func LoadMidtransConfig() {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
}
