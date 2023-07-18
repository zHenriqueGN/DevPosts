package routes

import (
	"api/internal/controller"

	"github.com/gofiber/fiber/v2"
)

// LoginRoute is a route to login functionality
var LoginRoute = Route{
	URI:          "/login",
	Method:       fiber.MethodPost,
	Func:         controller.Login,
	AuthRequired: true,
}
