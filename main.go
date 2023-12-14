package main

import (
	"latihan1/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.InitializeRoutes()

	//membuat logging
	log.Printf("Server started on port 8080")

	// memulai server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
