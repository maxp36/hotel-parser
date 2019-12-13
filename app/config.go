package app

import (
	"log"

	"github.com/joho/godotenv"
)

// InitConfig inits .env configuration
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
