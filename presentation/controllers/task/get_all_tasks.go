package task

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang-api-clean-architecture/core/contracts/task"
)

type Handler struct {
	getAllTasksUsecase task.GetAllTasksUsecase
	validator          *validator.Validate
}

func NewHandler(getAllTasksUsecase task.GetAllTasksUsecase) *GetAllTasksController {
	return &GetAllTasksController{
		getAllTasksUsecase: getAllTasksUsecase,
		validator:          validator.New(),
	}
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Get all tasks registered
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal error"
// @Router /tasks [get]
func (ctl *GetAllTasksController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := ctl.getAllTasksUsecase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
