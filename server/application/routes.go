package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"tema1/handler"
	"tema1/utils"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		
	})

	// Exercitiul 1
	router.Post("/concatByPosition",  handler.GenericHandler(utils.ConcatByPosition))
	
	return router
}

