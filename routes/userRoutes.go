package routes

import (
	"latihan1/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	//Routes Handling GET
	router.HandleFunc("/users/", handlers.GetUser).Methods("GET")

	//Routes Handling POST
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	return router

}
