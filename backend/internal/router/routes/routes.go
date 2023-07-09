package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Representes all the API routes
type Route struct {
	URI          string
	Method       string
	Func         func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

// Config register all the routes in the router
func Config(r *mux.Router) *mux.Router {
	routes := UserRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
