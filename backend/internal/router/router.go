package router

import (
	"api/internal/router/routes"

	"github.com/gofiber/fiber/v2"
)

// Generate generates the router
func Generate() (app *fiber.App) {
	app = fiber.New()
	app = routes.Config(app)
	return
}
