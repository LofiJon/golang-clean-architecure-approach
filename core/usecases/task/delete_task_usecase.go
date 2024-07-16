package task

import (
	"golang-api-clean-architecture/core/repositories"
)

type DeleteTaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewDeleteTaskUsecase(repo repositories.TaskRepository) *DeleteTaskUsecaseImpl {
	return &DeleteTaskUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *DeleteTaskUsecaseImpl) Execute(id string) error {
	return u.taskRepository.Delete(id)
}
