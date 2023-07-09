package routes

import (
	"api/internal/controller"
	"net/http"
)

// UserRoutes is a slice of Routes
var UserRoutes = []Route{
	{
		URI:          "/api/users",
		Method:       http.MethodPost,
		Func:         controller.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/api/users",
		Method:       http.MethodGet,
		Func:         controller.FetchUsers,
		AuthRequired: false,
	},
	{
		URI:          "/api/users/{id}",
		Method:       http.MethodGet,
		Func:         controller.FetchUser,
		AuthRequired: false,
	},
	{
		URI:          "/api/users/{id}",
		Method:       http.MethodPut,
		Func:         controller.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/api/users/{id}",
		Method:       http.MethodDelete,
		Func:         controller.DeleteUser,
		AuthRequired: false,
	},
}
