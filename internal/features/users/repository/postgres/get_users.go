package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/emp2ty0/golang-todoapp/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.poll.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	LIMIT $1
	OFFSET $2
	ORDER BY id ASC
	`

	rows, err := r.poll.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("select users:%w", err)
	}
	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel

		err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.PhoneNumber,
		)
		if err != nil {
			return nil, fmt.Errorf("scan users: %w", err)
		}

		userModels = append(userModels, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rowa: %w", err)
	}

	var userDomains []domain.User

	for _, user := range userModels {
		userDomain := domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)

		userDomains = append(userDomains, userDomain)
	}

	return userDomains, nil

}
