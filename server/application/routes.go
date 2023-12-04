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

	// Exercitiul 2
	router.Post("/numarPatratePerfecte",  handler.NumarPatratePerfecteHandler)

	// Exercitiul 3
	router.Post("/sumaInverselor", handler.SumaInverselorHandler)

	// Exercitiul 5
	router.Post("/conversieNumereBinare", handler.ConversieNumereBinareHandler)

	// Exercitiul 8
	router.Post("/numarCifreNumerelorPrime", handler.NumarCifreNumerelorPrimeHandler)

	// Exercitiul 9
	router.Post("/numarCuvinteVocalePare", handler.NumarCuvinteVocalePareHandler)

	// Exercitiul 10
	router.Post("/cmmdcNumere", handler.CMMDCNumereHandler)

	// Exercitiul 12
	router.Post("/sumaDublarePrimaCifra", handler.SumaDublarePrimaCifraHandler)

	// Exercitiul 14
	router.Post("/filtreazaParole", handler.FiltreazaParoleHandler)

	// Exercitiul 15
	router.Post("/genereazaParole", handler.GenereazaParoleHandler)

	
	return router
}

