package repositories

import "golang-api-clean-architecture/core/models"

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByID(id string) (models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	Count() (int64, error)
	GetPaged(page, size int) ([]models.User, error)
}
