package repositories

import "golang-api-clean-architecture/core/models"

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id string) (models.Task, error)
	GetAll() ([]models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
	Count() (int64, error)
	GetPaged(page, perPage int) ([]models.Task, error)
}
