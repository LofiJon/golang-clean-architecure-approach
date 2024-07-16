package task

import "golang-api-clean-architecture/core/models"

type GetByIdTask interface {
	Execute(id string) (models.Task, error)
}
