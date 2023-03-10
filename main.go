package main

import (
	"deweTourBE/database"
	"os"
	// "deweTourBE/handlers"
	"deweTourBE/pkg/mysql"
	"deweTourBE/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
//incometrip
	errNev := godotenv.Load()
	if errNev != nil {
		fmt.Println(errNev)
	}

	mysql.DatabaseInit()
	database.RunMigration()
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	AllowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	AllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")
	fmt.Println("server running localhost:" + port)
	// http.ListenAndServe("localhost:5000", r)

	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
