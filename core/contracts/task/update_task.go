package task

import "golang-api-clean-architecture/core/requests"

type UpdateTask interface {
	Execute(id string, task *requests.TaskRequest)
}
