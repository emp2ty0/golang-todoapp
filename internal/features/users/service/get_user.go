package users_service

import (
	"context"
	"fmt"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
)

func (u *UserService) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	user, err := u.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
