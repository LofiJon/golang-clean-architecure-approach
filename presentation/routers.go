package presentation

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang-api-clean-architecture/presentation/task"
)

func RegisterRoutes(r *mux.Router, handler *Handler) {
	sub := r.PathPrefix("/tasks").Subrouter()

	sub.HandleFunc("", handler.Create).Methods("POST")
	sub.HandleFunc("", handler.GetAll).Methods("GET")
	sub.HandleFunc("/paged", handler.GetPaged).Methods("GET")
	sub.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	sub.HandleFunc("/{id}", handler.Update).Methods("PUT")
	sub.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}