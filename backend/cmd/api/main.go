package main

import (
	"api/internal/config"
	"api/internal/router"
	"fmt"
)

func init() {
	config.LoadEnv()
}

func main() {
	app := router.Generate()
	app.Listen(fmt.Sprintf(":%s", config.APIPort))
}
