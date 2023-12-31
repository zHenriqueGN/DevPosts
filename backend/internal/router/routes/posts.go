package routes

import (
	"api/internal/controller"

	"github.com/gofiber/fiber/v2"
)

var PostRoutes = []Route{
	{
		URI:          "/",
		Method:       fiber.MethodPost,
		Func:         controller.CreatePost,
		AuthRequired: true,
	},
	{
		URI:          "/",
		Method:       fiber.MethodGet,
		Func:         controller.FetchPosts,
		AuthRequired: true,
	},
	{
		URI:          "/:id",
		Method:       fiber.MethodGet,
		Func:         controller.FetchPost,
		AuthRequired: true,
	},
	{
		URI:          "/:id",
		Method:       fiber.MethodPut,
		Func:         controller.UpdatePost,
		AuthRequired: true,
	},
	{
		URI:          "/:id",
		Method:       fiber.MethodDelete,
		Func:         controller.DeletePost,
		AuthRequired: true,
	},
	{
		URI:          "/user_posts/:id",
		Method:       fiber.MethodGet,
		Func:         controller.FetchPostsByUser,
		AuthRequired: true,
	},
	{
		URI:          "/:id/like",
		Method:       fiber.MethodPost,
		Func:         controller.LikePost,
		AuthRequired: true,
	},
	{
		URI:          "/:id/unlike",
		Method:       fiber.MethodPost,
		Func:         controller.UnlikePost,
		AuthRequired: true,
	},
}
