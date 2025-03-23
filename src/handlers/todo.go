package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	db "go-cc/src/database"
	"go-cc/src/models"
	"net/http"
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	db.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = uuid.NewString()
	db.DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo models.Todo
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedTodo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)

	var todo models.Todo
	if err := db.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	todo.Title = updatedTodo.Title
	todo.Done = updatedTodo.Done
	db.DB.Save(&todo)

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := db.DB.Where("id=?", id).Delete(&models.Todo{}).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Deleted successfully!"))
}
