package routes

import (
	"github.com/gorilla/mux"
	"../controllers"
)

// RegisterRoutes registers the routes for the application
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
}
