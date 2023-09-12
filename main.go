package main

import (
	"log"
	"net/http"
	"os"
	"github.com/hng/routes"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT") // Use the PORT environment variable

	if port == "" {
		port = "10000" // Set a default port (e.g., 10000) if PORT is not set
	}

	r := mux.NewRouter()
	routes.RegisterPersonRoutes(r)
	http.Handle("/", r)

	addr := ":" + port
	log.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
