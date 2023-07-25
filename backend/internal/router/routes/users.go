package routes

import (
	"api/internal/controller"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes is a slice of Routes to user functionalities
var UserRoutes = []Route{
	{
		URI:          "/",
		Method:       fiber.MethodPost,
		Func:         controller.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/",
		Method:       fiber.MethodGet,
		Func:         controller.FetchUsers,
		AuthRequired: true,
	},
	{
		URI:          "/:id",
		Method:       fiber.MethodGet,
		Func:         controller.FetchUser,
		AuthRequired: true,
	},
	{
		URI:          "/:id",
		Method:       fiber.MethodPut,
		Func:         controller.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/:id",
		Method:       fiber.MethodDelete,
		Func:         controller.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/:id/follow",
		Method:       fiber.MethodPost,
		Func:         controller.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/:id/unfollow",
		Method:       fiber.MethodPost,
		Func:         controller.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/:id/followers",
		Method:       fiber.MethodGet,
		Func:         controller.GetFollowers,
		AuthRequired: true,
	},
	{
		URI:          "/:id/followings",
		Method:       fiber.MethodGet,
		Func:         controller.GetFollowings,
		AuthRequired: true,
	},
}
