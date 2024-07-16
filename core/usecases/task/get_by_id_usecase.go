package task

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
)

type GetByIdTaskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewGetByIdTaskUsecase(repo repositories.TaskRepository) *GetByIdTaskUsecase {
	return &GetByIdTaskUsecase{
		taskRepository: repo,
	}
}

func (u *GetByIdTaskUsecase) Execute(id string) (models.Task, error) {
	return u.taskRepository.GetByID(id)
}
