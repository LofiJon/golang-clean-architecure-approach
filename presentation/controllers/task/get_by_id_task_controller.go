package task

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang-api-clean-architecture/core/contracts/task"
)

type GetByIdTaskController struct {
	getByIdTaskUsecase task.GetByIdTask
	validator          *validator.Validate
}

func NewGetByIdTaskController(
	getByIdTaskUsecase task.GetByIdTask) *GetByIdTaskController {
	return &GetByIdTaskController{
		getByIdTaskUsecase: getByIdTaskUsecase,
		validator:          validator.New(),
	}
}

// GetTask godoc
// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal error"
// @Router /tasks/{id} [get]
func (c *GetByIdTaskController) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task, err := c.getByIdTaskUsecase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}
