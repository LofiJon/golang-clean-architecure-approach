package users

import "golang-api-clean-architecture/core/repositories"

type CreateUserUsecase struct {
	userRepository repositories.UserRepository
}
