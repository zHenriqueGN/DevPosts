package main

import (
	"api/internal/config"
	"api/internal/router"
	"log"
	"net/http"
)

func init() {
	config.LoadEnv()
}

func main() {
	log.Println("Starting API...")
	r := router.Generate()
	log.Printf("Listening on http://localhost:%s", config.APIPort)
	log.Fatal(http.ListenAndServe(":8000", r))
}
