package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "golang-api-clean-architecture/docs" // Import the generated docs
	injection "golang-api-clean-architecture/infra/dependency_injection"
)

// @title Golang Clean Architecture Example
// @version 1.0
// @description This API is an approach of clean architecture using Golang made by Jonathan Malagueta
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth
func main() {
	// Initialize router
	router := injection.InitializeApp()

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Starting server at port 8080...")
	log.Println("Swagger at address: http://localhost:8080/swagger/index.html")
	http.ListenAndServe(":8080", router)
}
