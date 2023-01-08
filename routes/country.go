package routes

import (
	"deweTourBE/handlers"
	"deweTourBE/pkg/mysql"
	"deweTourBE/repositories"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(CountryRepository)

	r.HandleFunc("/country", h.FindCountry).Methods("GET")
	r.HandleFunc("/country/{id}", h.GetCountry).Methods("GET")
	r.HandleFunc("/country", h.CreateCountry).Methods("POST")
	r.HandleFunc("/country/{id}", h.UpdateCountry).Methods("PATCH")
	r.HandleFunc("/country/{id}", h.DeleteCountry).Methods("DELETE")
}
