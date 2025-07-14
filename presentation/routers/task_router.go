package routers

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang-api-clean-architecture/presentation/task"
)

func NewTaskRouter(handler *task.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", handler.Create).Methods("POST")
	router.HandleFunc("/tasks/paged", handler.GetPaged).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/tasks", handler.GetAll).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.Delete).Methods("DELETE")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}