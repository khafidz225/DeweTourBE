package routes

import (
	"deweTourBE/handlers"
	"deweTourBE/pkg/middleware"
	"deweTourBE/pkg/mysql"
	"deweTourBE/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(TripRepository)

	r.HandleFunc("/trip", h.FindTrip).Methods("GET")
	r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip))).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTrip))).Methods("PATCH")
	// r.HandleFunc("/trip/{id}", h.UpdateTrip).Methods("PATCH")
	r.HandleFunc("/trip/{id}", h.DeleteTrip).Methods("DELETE")
}
