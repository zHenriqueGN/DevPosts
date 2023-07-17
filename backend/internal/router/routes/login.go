package routes

import (
	"api/internal/controller"

	"github.com/gofiber/fiber/v2"
)

var LoginRoute = Route{
	URI:          "/login",
	Method:       fiber.MethodPost,
	Func:         controller.Login,
	AuthRequired: true,
}
