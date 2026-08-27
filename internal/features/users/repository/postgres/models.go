package users_postgres_repository

import "github.com/emp2ty0/golang-todoapp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func UserDomainsFromModels(users []UserModel) []domain.User {

	var userDomains = make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.FullName,
			user.PhoneNumber,
		)
	}

	return userDomains

}
