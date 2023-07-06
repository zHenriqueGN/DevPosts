package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBConn  string
	APIPort string
)

// LoadEnv loads the environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}
	DBConn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	APIPort = os.Getenv("API_PORT")
}
