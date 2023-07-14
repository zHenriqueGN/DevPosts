package router

import (
	"api/internal/router/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Generate generates the router
func Generate() (app *fiber.App) {
	app = fiber.New()
	app.Use(logger.New())
	app = routes.Config(app)
	return
}
