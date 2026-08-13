package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/emp2ty0/golang-todoapp/internal/core/errors"
)

type User struct {
	Id          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	full_name string,
	phone_number *string,
) User {
	return User{
		Id:          id,
		Version:     version,
		FullName:    full_name,
		PhoneNumber: phone_number,
	}
}

func NewUserUnitialized(
	full_name string,
	phone_number *string,
) User {
	return NewUser(UninitializedID, UninitializedVersion, full_name, phone_number)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("ivalid `Full Name` len: %d %w", fullNameLength, core_errors.ErrInvalidArgument)
	}

	if u.PhoneNumber != nil {
		phoneNumber := len([]rune(*u.PhoneNumber))
		if phoneNumber < 10 || phoneNumber > 15 {
			return fmt.Errorf("ivalid `Phone Number` len %d %w", phoneNumber, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("ivalid `Phone Number` len %d %w", phoneNumber, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}
