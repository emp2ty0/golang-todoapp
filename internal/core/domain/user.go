package domain

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
