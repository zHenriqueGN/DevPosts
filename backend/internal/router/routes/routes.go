package routes

import (
	"api/internal/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Representes all the API routes
type Route struct {
	URI          string
	Method       string
	Func         func(*fiber.Ctx) error
	AuthRequired bool
}

// Config register all the routes in the router
// and applies the middlewares
func Config(app *fiber.App) *fiber.App {
	app.Use(logger.New())

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Add(LoginRoute.Method, LoginRoute.URI, LoginRoute.Func)

	users := api.Group("/users")
	for _, route := range UserRoutes {
		if route.AuthRequired {
			users.Add(route.Method, route.URI, jwtware.New(
				jwtware.Config{
					SigningKey:     jwtware.SigningKey{Key: []byte(config.SecretKey)},
					SuccessHandler: route.Func,
				}))
		} else {
			users.Add(route.Method, route.URI, route.Func)
		}
	}

	posts := api.Group("/posts")
	for _, route := range PostRoutes {
		if route.AuthRequired {
			posts.Add(route.Method, route.URI, jwtware.New(
				jwtware.Config{
					SigningKey:     jwtware.SigningKey{Key: []byte(config.SecretKey)},
					SuccessHandler: route.Func,
				},
			))
		} else {
			posts.Add(route.Method, route.URI, route.Func)
		}
	}

	return app
}
