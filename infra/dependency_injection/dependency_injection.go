package dependency_injection

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	repository "golang-api-clean-architecture/core/repositories"
	taskUsecases "golang-api-clean-architecture/core/usecases/task"
	"golang-api-clean-architecture/infra/databse"
	"golang-api-clean-architecture/infra/repositories"
	taskController "golang-api-clean-architecture/presentation/controllers/task"
	"golang-api-clean-architecture/presentation/routers"
	"gorm.io/gorm"
)

// InitializeDatabase sets up the database connection
func InitializeDatabase() *gorm.DB {
	return database.InitPostgres()
}

// InitializeRepositories sets up the repositories
func InitializeRepositories(db *gorm.DB) repository.TaskRepository {
	return repositories.NewTaskRepository(db)
}

// InitializeUseCases sets up the use cases
func InitializeUseCases(taskRepo repository.TaskRepository) (
	*taskUsecases.CreateTaskUsecaseImpl,
	*taskUsecases.GetByIdTaskUsecaseImpl,
	*taskUsecases.GetAllTasksUsecaseImpl,
	*taskUsecases.UpdateTaskUsecaseImpl,
	*taskUsecases.DeleteTaskUsecaseImpl,
	*taskUsecases.PageableTaskUsecaseImpl,
) {
	return taskUsecases.NewCreateTaskUsecase(taskRepo),
		taskUsecases.NewGetByIdTaskUsecase(taskRepo),
		taskUsecases.NewGetAllTasksUsecase(taskRepo),
		taskUsecases.NewUpdateTaskUsecase(taskRepo),
		taskUsecases.NewDeleteTaskUsecase(taskRepo),
		taskUsecases.NewPageableTaskUsecase(taskRepo)
}

// InitializeControllers sets up the controllers
func InitializeControllers(
	createTaskUsecase *taskUsecases.CreateTaskUsecaseImpl,
	getByIdTaskUsecase *taskUsecases.GetByIdTaskUsecaseImpl,
	getAllTasksUsecase *taskUsecases.GetAllTasksUsecaseImpl,
	updateTaskUsecase *taskUsecases.UpdateTaskUsecaseImpl,
	deleteTaskUsecase *taskUsecases.DeleteTaskUsecaseImpl,
	pageableTaskUsecase *taskUsecases.PageableTaskUsecaseImpl,
) (
	*taskController.CreateTaskController,
	*taskController.GetByIdTaskController,
	*taskController.GetAllTasksController,
	*taskController.UpdateTaskController,
	*taskController.DeleteTaskController,
	*taskController.PageableTaskController,
) {
	return taskController.NewCreateTaskController(createTaskUsecase),
		taskController.NewGetByIdTaskController(getByIdTaskUsecase),
		taskController.NewGetAllTasksController(getAllTasksUsecase),
		taskController.NewUpdateTaskController(updateTaskUsecase),
		taskController.NewDeleteTaskController(deleteTaskUsecase),
		taskController.NewPageableTaskController(pageableTaskUsecase)
}

// InitializeRouter sets up the router
func InitializeRouter(
	createTaskController *taskController.CreateTaskController,
	getByIdTaskController *taskController.GetByIdTaskController,
	getAllTasksController *taskController.GetAllTasksController,
	updateTaskController *taskController.UpdateTaskController,
	deleteTaskController *taskController.DeleteTaskController,
	pageableTaskController *taskController.PageableTaskController,
) *mux.Router {
	router := routers.NewTaskRouter(
		createTaskController,
		getByIdTaskController,
		getAllTasksController,
		updateTaskController,
		pageableTaskController,
		deleteTaskController,
	)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return router
}

// InitializeApp sets up the entire application
func InitializeApp() *mux.Router {
	// Configure database
	db := InitializeDatabase()

	// Initialize repositories
	taskRepo := InitializeRepositories(db)

	// Initialize use cases
	createTaskUsecase, getByIdTaskUsecase, getAllTasksUsecase, updateTaskUsecase, deleteTaskUsecase, pageableTaskUsecase := InitializeUseCases(taskRepo)

	// Initialize controllers
	createTaskController, getByIdTaskController, getAllTasksController, updateTaskController, deleteTaskController, pageableTaskController := InitializeControllers(
		createTaskUsecase,
		getByIdTaskUsecase,
		getAllTasksUsecase,
		updateTaskUsecase,
		deleteTaskUsecase,
		pageableTaskUsecase,
	)

	// Initialize router
	router := InitializeRouter(
		createTaskController,
		getByIdTaskController,
		getAllTasksController,
		updateTaskController,
		deleteTaskController,
		pageableTaskController,
	)

	return router
}
