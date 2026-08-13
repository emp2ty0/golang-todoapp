package users_postgres_repository

import core_postgres_pool "github.com/emp2ty0/golang-todoapp/internal/core/repository/postgres/pool"

type UsersRepository struct {
	poll core_postgres_pool.Poll
}

func NewUsersRepository(
	poll core_postgres_pool.Poll,
) *UsersRepository {
	return &UsersRepository{
		poll: poll,
	}
}
