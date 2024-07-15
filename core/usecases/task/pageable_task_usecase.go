package task

import (
	"golang-api-clean-architecture/core/contracts/task"
	"golang-api-clean-architecture/core/dtos"
	"golang-api-clean-architecture/core/repositories"
	"golang-api-clean-architecture/core/requests"
)

type pageableTaskUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewPageableTaskUsecase(repo repositories.TaskRepository) task.PageableTask {
	return &pageableTaskUsecase{
		taskRepository: repo,
	}
}

func (u *pageableTaskUsecase) GetTasks(pageRequest requests.PageRequest) (dtos.PageableDto, error) {
	totalItems, err := u.taskRepository.Count()
	if err != nil {
		return dtos.PageableDto{}, err
	}

	tasks, err := u.taskRepository.Pageable(pageRequest.Page, pageRequest.Size)
	if err != nil {
		return dtos.PageableDto{}, err
	}

	totalPages := int(totalItems) / pageRequest.Size
	if int(totalItems)%pageRequest.Size != 0 {
		totalPages++
	}

	return dtos.PageableDto{
		Page:       pageRequest.Page,
		Size:       pageRequest.Size,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Items:      tasks,
	}, nil
}
