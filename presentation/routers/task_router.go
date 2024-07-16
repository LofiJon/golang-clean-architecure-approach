package routers

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang-api-clean-architecture/presentation/controllers/task"
)

func NewTaskRouter(
	createTaskController *task.CreateTaskController,
	getByIdTaskController *task.GetByIdTaskController,
	getAllTasksController *task.GetAllTasksController,
	updateTaskController *task.UpdateTaskController,
	pagedTaskController *task.PageableTaskController,
	deleteTaskController *task.DeleteTaskController,
) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", createTaskController.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/paged", pagedTaskController.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getByIdTaskController.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks", getAllTasksController.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTaskController.UpdateTaskByID).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTaskController.DeleteTaskByID).Methods("DELETE")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
