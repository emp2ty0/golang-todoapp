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

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
}

func NewUserService(usersRepository UsersRepositoryInterface) *UserService {
	return &UserService{
		usersRepository: usersRepository,
	}
}
