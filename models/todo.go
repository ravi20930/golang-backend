package models

// Todo represents a todo item
type Todo struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
