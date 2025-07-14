package dependency_injection

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang-api-clean-architecture/core/contracts/task"
	taskUsecases "golang-api-clean-architecture/core/usecases/task"
	"golang-api-clean-architecture/infra/databse"
	"golang-api-clean-architecture/infra/repositories"
	"golang-api-clean-architecture/presentation/task"
)

func InitializeApp() *mux.Router {
	db := database.InitPostgres()

	taskRepo := repositories.NewTaskRepository(db)

	
	createTask := taskUsecases.NewCreateTaskUsecase(taskRepo)
	getByIdTask := taskUsecases.NewGetByIdTaskUsecase(taskRepo)
	getAllTasks := taskUsecases.NewGetAllTasksUsecase(taskRepo)
	updateTask := taskUsecases.NewUpdateTaskUsecase(taskRepo)
	deleteTask := taskUsecases.NewDeleteTaskUsecase(taskRepo)
	pageableTask := taskUsecases.NewPageableTaskUsecase(taskRepo)

	taskHandler := task.NewHandler(
		createTask,
		getByIdTask,
		getAllTasks,
		updateTask,
		deleteTask,
		pageableTask,
	)

	router := mux.NewRouter().StrictSlash(true)
	task.RegisterRoutes(router, taskHandler)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}

func NewHandler(
	create task.CreateTask,
	getByID task.GetByIdTask,
	getAll task.GetAllTasks,
	update task.UpdateTask,
	delete task.DeleteTask,
	pageable task.PageableTask,
) *Handler {
	return &Handler{
		createTaskUsecase:  create,
		getByIdTaskUsecase: getByID,
		getAllTasksUsecase: getAll,
		updateTaskUsecase:  update,
		deleteTaskUsecase:  delete,
		pageableTaskUsecase: pageable,
		validator:           validator.New(),
	}
}