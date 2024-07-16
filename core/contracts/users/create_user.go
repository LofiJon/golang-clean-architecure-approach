package users

import "golang-api-clean-architecture/core/requests"

type CreateUser interface {
	Execute(user *requests.UserRequest)
}
