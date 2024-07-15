package task

import (
	"golang-api-clean-architecture/core/dtos"
	"golang-api-clean-architecture/core/requests"
)

type PageableTask interface {
	GetTasks(pageRequest requests.PageRequest) (dtos.PageableDto, error)
}
