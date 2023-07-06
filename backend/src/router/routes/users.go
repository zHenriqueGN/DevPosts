package routes

import "net/http"

var UserRoutes = []Route{
	{
		URI:    "/api/users",
		Method: http.MethodPost,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/api/users",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/api/users/{id}",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/api/users/{id}",
		Method: http.MethodPut,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/api/users/{id}",
		Method: http.MethodDelete,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
}
