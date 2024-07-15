package task

import "golang-api-clean-architecture/core/requests"

type CreateTask interface {
	Execute(task *requests.TaskRequest) error
}
