package dependency_injection

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	taskContracts "golang-api-clean-architecture/core/contracts/task"
	taskUsecases "golang-api-clean-architecture/core/usecases/task"
	"golang-api-clean-architecture/infra/databse"
	"golang-api-clean-architecture/infra/repositories"
	"golang-api-clean-architecture/presentation/task"
)

func InitializeApp() *mux.Router {
	// Conexão com o banco
	db := database.InitPostgres()

	// Repository
	taskRepo := repositories.NewTaskRepository(db)

	// UseCases (injeção por interface)
	var (
		createTask    taskContracts.CreateTask = taskUsecases.NewCreateTaskUsecase(taskRepo)
		getByIdTask   taskContracts.GetByIdTask = taskUsecases.NewGetByIdTaskUsecase(taskRepo)
		getAllTasks   taskContracts.GetAllTasks = taskUsecases.NewGetAllTasksUsecase(taskRepo)
		updateTask    taskContracts.UpdateTask = taskUsecases.NewUpdateTaskUsecase(taskRepo)
		deleteTask    taskContracts.DeleteTask = taskUsecases.NewDeleteTaskUsecase(taskRepo)
		pageableTask  taskContracts.PageableTask = taskUsecases.NewPageableTaskUsecase(taskRepo)
	)

	// Handler unificado
	taskHandler := task.NewHandler(
		createTask,
		getByIdTask,
		getAllTasks,
		updateTask,
		deleteTask,
		pageableTask,
	)

	// Router
	router := mux.NewRouter().StrictSlash(true)
	task.RegisterRoutes(router, taskHandler)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
