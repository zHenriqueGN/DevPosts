package router

import (
	"api/internal/router/routes"
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
)

// Generate generates the router
func Generate() (r *mux.Router) {
	r = mux.NewRouter()
	r = routes.Config(r)
	r.Use(LoggingMiddleware)
	return
}

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(h, w, r)
		log.Printf(
			"%s %s %d %v %d",
			r.Method,
			r.URL,
			m.Code,
			m.Duration,
			m.Written,
		)
	})
}
