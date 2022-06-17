package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration model
type Configuration struct {
	Port             string
	ConnectionString string
	Database         string
	Collection       string
}

// GetConfiguration method basically populate configuration information from .env and return Configuration model
func GetConfiguration() Configuration {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configuration := Configuration{
		os.Getenv("port"),
		os.Getenv("connection"),
		os.Getenv("database"),
		os.Getenv("collection"),
	}

	return configuration
}
