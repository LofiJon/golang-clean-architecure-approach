package task

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang-api-clean-architecture/core/contracts/task"
	request "golang-api-clean-architecture/core/requests"
	"net/http"
)

type CreateTaskController struct {
	createTaskUsecase task.CreateTask
	validator         *validator.Validate
}

func NewCreateTaskController(createTaskUsecase task.CreateTask) *CreateTaskController {
	return &CreateTaskController{
		createTaskUsecase: createTaskUsecase,
		validator:         validator.New(),
	}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body request.TaskRequest true "Task to create"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal error"
// @Router /tasks [post]
func (c *CreateTaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req request.TaskRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	err := c.validator.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.createTaskUsecase.Execute(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
