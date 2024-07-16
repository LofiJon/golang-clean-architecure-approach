package task

import (
	"golang-api-clean-architecture/core/contracts/task"
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
)

var _ task.GetAllTasksUsecase = &GetAllTasksUsecaseImpl{}

type GetAllTasksUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewGetAllTasksUsecase(repo repositories.TaskRepository) *GetAllTasksUsecaseImpl {
	return &GetAllTasksUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *GetAllTasksUsecaseImpl) Execute() ([]models.Task, error) {
	return u.taskRepository.GetAll()
}
