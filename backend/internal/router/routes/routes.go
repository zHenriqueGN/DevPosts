package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Representes all the API routes
type Route struct {
	URI          string
	Method       string
	Func         func(*fiber.Ctx) error
	AuthRequired bool
}

// Config register all the routes in the router
func Config(app *fiber.App) *fiber.App {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Add(LoginRoute.Method, LoginRoute.URI, LoginRoute.Func)

	users := api.Group("/users")
	for _, route := range UserRoutes {
		users.Add(route.Method, route.URI, route.Func)
	}

	return app
}
