package task

import (
	"golang-api-clean-architecture/core/contracts/task"
	"golang-api-clean-architecture/core/repositories"
)

type deleteTaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewDeleteTaskUsecase(repo repositories.TaskRepository) task.DeleteTask {
	return &deleteTaskUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *deleteTaskUsecaseImpl) Execute(id string) error {
	return u.taskRepository.Delete(id)
}
