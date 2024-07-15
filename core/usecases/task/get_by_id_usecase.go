package task

import "golang-api-clean-architecture/core/repositories"

type FindByIdUsecaseImpl struct {
	taskRepository repositories.TaskRepository
}

func NewGetByIdTaskUsecase(repo repositories.TaskRepository) *FindByIdUsecaseImpl {
	return &FindByIdUsecaseImpl{repo}
}

func (uc *FindByIdUsecaseImpl) FindById(id string) (interface{}, error) {
	return uc.taskRepository.GetByID(id)
}
