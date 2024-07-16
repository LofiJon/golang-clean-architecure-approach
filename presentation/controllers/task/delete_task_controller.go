package task

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	usecase "golang-api-clean-architecture/core/contracts/task"
	"net/http"
)

type DeleteTaskController struct {
	deleteTaskUsecase usecase.DeleteTask
	validator         *validator.Validate
}

func NewDeleteTaskController(
	deleteTaskUsecase usecase.DeleteTask) *DeleteTaskController {
	return &DeleteTaskController{
		deleteTaskUsecase: deleteTaskUsecase,
		validator:         validator.New(),
	}
}

// DeleteTask godoc
// @Summary Delete task by id
// @Description Delete task by id
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Internal error"
// @Router /tasks/{id} [delete]
func (c *DeleteTaskController) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	task := c.deleteTaskUsecase.Execute(id)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(task)
}
