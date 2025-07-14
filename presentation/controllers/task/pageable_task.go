package task

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"golang-api-clean-architecture/core/contracts/task"
	"golang-api-clean-architecture/core/requests"
)

// PageableTaskController handles paginated task requests
type Handler struct {
	pageableTaskUsecase task.PageableTask
	validator           *validator.Validate
}

// NewPageableTaskController creates a new PageableTaskController
func NewHandler(pageableTaskUsecase task.PageableTask) *PageableTaskController {
	return &PageableTaskController{
		pageableTaskUsecase: pageableTaskUsecase,
		validator:           validator.New(),
	}
}

// GetTasks godoc
// @Summary Get paginated tasks
// @Description Get paginated tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param per_page query int true "Number of items per page"
// @Success 200 {object} dtos.PageableDto
// @Failure 400 {object} dtos.HTTPError
// @Failure 500 {object} dtos.HTTPError
// @Router /tasks/paged [get]
func (c *PageableTaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		http.Error(w, "Invalid per_page parameter", http.StatusBadRequest)
		return
	}

	pageRequest := requests.PageRequest{
		Page: page,
		Size: perPage,
	}

	err = c.validator.Struct(pageRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := c.pageableTaskUsecase.GetTasks(pageRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
