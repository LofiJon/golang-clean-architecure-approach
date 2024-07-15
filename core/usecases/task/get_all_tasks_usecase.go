package task

import "golang-api-clean-architecture/core/repositories"

type GetAllTasksUsecase struct {
	taskRepository repositories.TaskRepository
}

func NewGetAllTaskUsecase(repo repositories.TaskRepository) *GetAllTasksUsecase {
	return &GetAllTasksUsecase{repo}
}

func (uc *GetAllTasksUsecase) GetAll() (interface{}, error) {
	return uc.taskRepository.GetAll()
}
