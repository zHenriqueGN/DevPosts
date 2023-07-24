package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// DBConn is the connection string to database
	DBConn string

	// APIPort is the port on the API will be running
	APIPort string

	// SecretKey is the key used to sign the web-token
	SecretKey string
)

// LoadEnv loads the environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}
	DBConn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	APIPort = os.Getenv("API_PORT")
	SecretKey = os.Getenv("SECRET_KEY")
}
