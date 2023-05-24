package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"./routes"
)

func main() {
	// Initialize router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
