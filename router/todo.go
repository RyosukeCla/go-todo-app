package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func addTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Added"))
}

func TodoRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/add", addTodo)
	return router
}
