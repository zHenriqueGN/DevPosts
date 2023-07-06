package router

import "github.com/gorilla/mux"

// Generate the router
func Generate() *mux.Router {
	return mux.NewRouter()
}
