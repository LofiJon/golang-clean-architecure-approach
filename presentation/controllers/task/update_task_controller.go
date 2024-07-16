package task

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang-api-clean-architecture/core/contracts/task"
	_ "golang-api-clean-architecture/core/dtos"
	"golang-api-clean-architecture/core/requests"
	"net/http"
)

// UpdateTaskController handles requests to update a task
type UpdateTaskController struct {
	updateTaskUsecase task.UpdateTask
	validator         *validator.Validate
}

// NewUpdateTaskController creates a new UpdateTaskController
func NewUpdateTaskController(updateTaskUsecase task.UpdateTask) *UpdateTaskController {
	return &UpdateTaskController{
		updateTaskUsecase: updateTaskUsecase,
		validator:         validator.New(),
	}
}

// UpdateTaskByID godoc
// @Summary Update a task by ID
// @Description Update a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body requests.TaskRequest true "Task data"
// @Success 200 {object} models.Task
// @Failure 400 {object} dtos.HTTPError
// @Failure 404 {object} dtos.HTTPError
// @Failure 500 {object} dtos.HTTPError
// @Router /tasks/{id} [put]
func (c *UpdateTaskController) UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req requests.TaskRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	err := c.validator.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.updateTaskUsecase.Execute(id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
