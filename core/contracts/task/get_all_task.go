package task

import "golang-api-clean-architecture/core/models"

type GetAllTasksUsecase interface {
	Execute() ([]models.Task, error)
}
