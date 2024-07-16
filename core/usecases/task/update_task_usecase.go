package task

import (
	"golang-api-clean-architecture/core/contracts/task"
	"golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/core/requests"
)

var _ task.UpdateTask = &UpdateTaskUsecaseImpl{}

type UpdateTaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewUpdateTaskUsecase(repo repositories.TaskRepository) *UpdateTaskUsecaseImpl {
	return &UpdateTaskUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *UpdateTaskUsecaseImpl) Execute(id string, taskRequest *requests.TaskRequest) error {
	task, err := u.taskRepository.GetByID(id)
	if err != nil {
		return err
	}

	task.Name = taskRequest.Name
	task.Done = taskRequest.Done

	return u.taskRepository.Update(&task)
}
