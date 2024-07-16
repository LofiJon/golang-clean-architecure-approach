package task

import (
	"golang-api-clean-architecture/core/models"
	"golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/core/requests"
)

type CreateTaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewCreateTaskUsecase(repo repositories.TaskRepository) *CreateTaskUsecaseImpl {
	return &CreateTaskUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *CreateTaskUsecaseImpl) Execute(taskRequest *requests.TaskRequest) error {
	taskModel := &models.Task{
		Name: taskRequest.Name,
		Done: taskRequest.Done,
	}
	return u.taskRepository.Create(taskModel)
}
