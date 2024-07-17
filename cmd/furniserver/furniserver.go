package main

import (
	"log"
	"net/http"

	"github.com/Dmitrygosu/furniture-rest-api/internal/app"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	app.RegisterRoutes(router)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
