package utils

import (
	"github.com/joho/godotenv"
	"log"
)

// TODO: call this function to load the environment variables
func LoadEnv() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
