package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate generates the router
func Generate() (r *mux.Router) {
	r = mux.NewRouter()
	r = routes.Config(r)
	return
}
