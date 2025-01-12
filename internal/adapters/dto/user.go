package dto

type PasswordStruct struct {
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type User struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	PasswordStruct
}
