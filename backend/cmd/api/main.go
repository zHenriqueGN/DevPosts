package main

import (
	"api/internal/config"
	"api/internal/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	app := router.Generate()
	app.Listen(":8000")
}
