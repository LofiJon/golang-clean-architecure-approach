package main

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Golang clean architecture example
// @version 1.0
// @description This api is an approach of clean architecture using golang made by Jonathan Malagueta
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter().StrictSlash(true)
	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Starting server at port 8080...")
	log.Println("Swagger at address: http://localhost:8080/swagger/index.html")
	http.ListenAndServe(":8080", router)
}
