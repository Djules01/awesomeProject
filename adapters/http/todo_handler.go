package httpadapter

import (
	"awesomeProject/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service domain.TodoService
}

func NewHandler(service domain.TodoService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var input domain.TodoList

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	todo, err := h.Service.CreateTodo(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) GetTodosByDate(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetTodoByDate()
	if err != nil {
		http.Error(w, "Erreur SQL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var input domain.TodoList

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	todo, ok, err := h.Service.UpdateTodo(id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !ok {
		http.Error(w, "Todo introuvable", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(todo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	ok, err := h.Service.DeleteTodo(id)
	if err != nil {
		http.Error(w, "Erreur SQL", http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Todo introuvable", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
