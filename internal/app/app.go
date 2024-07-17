package app

import (
	"github.com/gorilla/mux"
	"github.com/Dmitrygosu/furniture-rest-api/internal/service"
	"net/http"
)
// RegisterRoutes устанавлиает HTTP маршруты для приложения
func RegisterRoutes(router *mux.Router) {
	handler := service.NewService("database.json")

	router.HandleFunc("/furniture", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/furniture", handler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/furniture/{id:[0-9]+}", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/furniture/{id:[0-9]+}", handler.Update).Methods(http.MethodPut)
	router.HandleFunc("/furniture/{id:[0-9]+}", handler.Patch).Methods(http.MethodPatch)
	router.HandleFunc("/furniture/{id:[0-9]+}", handler.Delete).Methods(http.MethodDelete)
}
