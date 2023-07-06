package routes

import "net/http"

// Representes all the API routes
type Route struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, http.Request)
	AuthRequired bool
}
