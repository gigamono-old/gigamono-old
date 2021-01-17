package envs

import (
	"log"
	"github.com/joho/godotenv"
)

// LoadEnvFile loads .env file.
func LoadEnvFile() {
	err := godotenv.Load() // Loads .env file
	if err != nil {
		log.Println("Unable load environment variables:", err)
	}
}
