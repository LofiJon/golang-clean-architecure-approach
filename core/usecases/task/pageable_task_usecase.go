package task

import (
	"golang-api-clean-architecture/core/dtos"
	"golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/core/requests"
)

type PageableTaskUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewPageableTaskUsecase(repo repositories.TaskRepository) *PageableTaskUsecaseImpl {
	return &PageableTaskUsecaseImpl{
		taskRepository: repo,
	}
}

func (u *PageableTaskUsecaseImpl) GetTasks(pageRequest requests.PageRequest) (dtos.PageableDto, error) {
	tasks, err := u.taskRepository.GetPaged(pageRequest.Page, pageRequest.Size)
	if err != nil {
		return dtos.PageableDto{}, err
	}

	total, err := u.taskRepository.Count()
	if err != nil {
		return dtos.PageableDto{}, err
	}

	return dtos.PageableDto{
		Items: tasks,
		Size:  int(total),
		Page:  pageRequest.Page,
	}, nil
}
