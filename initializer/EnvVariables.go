package initializer

import (
	"github.com/joho/godotenv"

    "log"
)
func EnvVariables() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}