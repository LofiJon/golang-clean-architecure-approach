package task

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	usecase "golang-api-clean-architecture/core/contracts/task"
	request "golang-api-clean-architecture/core/requests"
)

type UpdateTaskController struct {
	updateTaskUsecase usecase.UpdateTask
	validator         *validator.Validate
}

func NewUpdateTaskController(updateTaskUsecase usecase.UpdateTask) *UpdateTaskController {
	return &UpdateTaskController{
		updateTaskUsecase: updateTaskUsecase,
		validator:         validator.New(),
	}
}

// UpdateTaskByID godoc
// @Summary Update a task
// @Description Update a task by ID
// @Tags tasks
// @Param id path string true "Task ID"
// @Param task body request.TaskRequest true "Task"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [put]
func (ctl *UpdateTaskController) UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req request.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctl.validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctl.updateTaskUsecase.Execute(id, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
