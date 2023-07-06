package main

import (
	"api/src/config"
	"api/src/router"
	"log"
	"net/http"
)

func init() {
	config.LoadEnv()
}

func main() {
	log.Println("Starting API...")
	r := router.Generate()
	log.Println("Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
