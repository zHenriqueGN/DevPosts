package router

import (
	"api/src/router/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Generate generates the router
func Generate() (r *mux.Router) {
	r = mux.NewRouter()
	r = routes.Config(r)
	r.Use(logsMiddleware)
	return
}

func logsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		res := &responseLogger{w, http.StatusOK}
		next.ServeHTTP(res, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s %s %d %v",
			r.Method,
			r.RequestURI,
			r.Proto,
			res.status,
			duration,
		)
	})
}

type responseLogger struct {
	http.ResponseWriter
	status int
}

func (r *responseLogger) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
