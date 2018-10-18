package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RyosukeCla/go-todo-app/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// Api
	r.Route("/api", func(r chi.Router) {
		r.Mount("/todo", router.TodoRouter())
	})

	// Render Client App
	r.Mount("/*", router.ClientAppRouter())

	port := 3000
	fmt.Println(fmt.Sprintf("Server listening on port %d", port))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
