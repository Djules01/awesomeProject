package httpadapter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handler *Handler, apiKey string) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(APIKeyMiddleware(apiKey))

	r.Get("/ToDoList", handler.Welcome)
	r.Post("/Creation", handler.CreateTodo)
	r.Get("/AfficherParDate", handler.GetTodosByDate)
	r.Put("/Modifier/{id}", handler.UpdateTodo)
	r.Delete("/Delete/{id}", handler.DeleteTodo)

	return r
}
