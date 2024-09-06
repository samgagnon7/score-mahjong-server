package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Mount("/books", ImageRoutes())

	http.ListenAndServe(":3000", r)
}

func ImageRoutes() chi.Router {
	r := chi.NewRouter()
	imageHandler := ImageHandler{} // Create an instance of ImageHandler
	r.Post("/", imageHandler.PostImage)
	return r
}
