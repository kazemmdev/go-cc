package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type Todo struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	DB.Where("").Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = uuid.NewString()
	DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo Todo
	if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)

	var todo Todo
	if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todo.Title = updatedTodo.Title
	todo.Done = updatedTodo.Done
	DB.Save(&todo)

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Deleted successfully!"))
}
