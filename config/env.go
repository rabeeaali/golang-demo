package config

import (
	"github.com/joho/godotenv"
	"log"
)

//LoadEnv loads environment variables from .env file
func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("unable to load .env file")
	}
}
