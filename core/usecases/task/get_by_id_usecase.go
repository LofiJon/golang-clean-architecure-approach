package task

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
)

type GetByIdTaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewGetByIdTaskUsecase(repo repositories.TaskRepository) *GetByIdTaskUsecaseImpl {
	return &GetByIdTaskUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *GetByIdTaskUsecaseImpl) Execute(id string) (models.Task, error) {
	return u.taskRepository.GetByID(id)
}
