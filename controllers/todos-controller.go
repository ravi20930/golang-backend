package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"../models"
)

var todos []models.Todo

// GetTodos returns all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo creates a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo updates an existing todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos[index].Title = item.Title
			todos[index].Completed = item.Completed
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

// DeleteTodo deletes a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
