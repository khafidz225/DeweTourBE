package routes

import (
	"deweTourBE/handlers"
	"deweTourBE/pkg/middleware"
	"deweTourBE/pkg/mysql"
	"deweTourBE/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	AuthRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(AuthRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	// r.HandleFunc("/login", h.Login).Methods("GET")
	r.HandleFunc("/login", h.Login).Methods("POST")

	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
