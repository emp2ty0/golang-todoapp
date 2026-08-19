package users_service

import (
	"context"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int,
) error {
	if err := s.usersRepository.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}
