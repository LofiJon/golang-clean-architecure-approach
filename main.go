package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	usecases "golang-api-clean-architecture/core/usecases/task"
	_ "golang-api-clean-architecture/docs" // Import the generated docs
	"golang-api-clean-architecture/infra/databse"
	"golang-api-clean-architecture/infra/repositories"
	"golang-api-clean-architecture/presentation/controllers/task"
	"golang-api-clean-architecture/presentation/routers"
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
	// Configure database
	db := database.InitPostgres()

	// Initialize repositories
	taskRepo := repositories.NewTaskRepository(db)

	// Initialize use cases
	createTaskUsecase := usecases.NewCreateTaskUsecase(taskRepo)
	getByIdTaskUsecase := usecases.NewGetByIdTaskUsecase(taskRepo)
	getAllTasksUsecase := usecases.NewGetAllTasksUsecase(taskRepo)
	updateTaskUsecase := usecases.NewUpdateTaskUsecase(taskRepo)
	pageableTaskUsecase := usecases.NewPageableTaskUsecase(taskRepo)
	deleteTaskUsecase := usecases.NewDeleteTaskUsecase(taskRepo)

	// Initialize controllers
	createTaskController := task.NewCreateTaskController(createTaskUsecase)
	getByIdTaskController := task.NewGetByIdTaskController(getByIdTaskUsecase)
	getAllTasksController := task.NewGetAllTasksController(getAllTasksUsecase)
	updateTaskController := task.NewUpdateTaskController(updateTaskUsecase)
	pageableTaskController := task.NewPageableTaskController(pageableTaskUsecase)

	deleteTaskController := task.NewDeleteTaskController(deleteTaskUsecase)

	// Initialize router
	router := routers.NewTaskRouter(createTaskController, getByIdTaskController, getAllTasksController, updateTaskController, pageableTaskController, deleteTaskController)

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Starting server at port 8080...")
	log.Println("Swagger at address: http://localhost:8080/swagger/index.html")
	http.ListenAndServe(":8080", router)
}
