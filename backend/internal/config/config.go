package config

import (
	"fmt"
	"os"
)

var (
	// DBConn is the connection string to database
	DBConn string

	// SecretKey is the key used to sign the web-token
	SecretKey string
)

// LoadEnv loads the environment variables
func LoadEnv() {
	DBConn = fmt.Sprintf(
		"host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	SecretKey = os.Getenv("SECRET_KEY")
}
