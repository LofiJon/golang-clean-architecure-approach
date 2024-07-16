package task

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/core/requests"
)

type CreateTaskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewCreateTaskUsecase(repo repositories.TaskRepository) *CreateTaskUsecase {
	return &CreateTaskUsecase{
		taskRepository: repo,
	}
}

func (u *CreateTaskUsecase) Execute(taskRequest *requests.TaskRequest) error {
	taskModel := &models.Task{
		Name: taskRequest.Name,
		Done: taskRequest.Done,
	}
	return u.taskRepository.Create(taskModel)
}
