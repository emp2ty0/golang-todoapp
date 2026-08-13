package users_service

import (
	"context"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
)

type UserService struct {
	usersRepository UsersRepositoryInterface
}

type UsersRepositoryInterface interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUserService(usersRepository UsersRepositoryInterface) *UserService {
	return &UserService{
		usersRepository: usersRepository,
	}
}
