package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id *int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.poll.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	WHERE id = $1
	`

	row := r.poll.QueryRow(ctx, query, *id)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(userModel.ID, userModel.Version, userModel.FullName, userModel.PhoneNumber)

	return userDomain, nil

}
