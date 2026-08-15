package users_service

import (
	"context"
	"fmt"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
	core_errors "github.com/emp2ty0/golang-todoapp/internal/core/errors"
)

func (u *UserService) GetUser(
	ctx context.Context,
	id *int,
) (domain.User, error) {
	if id != nil && *id <= 0 {
		return domain.User{}, fmt.Errorf("user id not be negative: %w", core_errors.ErrInvalidArgument)
	}

	user, err := u.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
