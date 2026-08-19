package users_service

import (
	"context"
	"fmt"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
	core_errors "github.com/emp2ty0/golang-todoapp/internal/core/errors"
)

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit not be negative %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset not be negative %w", core_errors.ErrInvalidArgument)
	}

	userDomains, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return userDomains, nil
}
