package userdata

import "github.com/go-playground/validator"

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (l *Login) Validate() error {
	validator := validator.New()
	return validator.Struct(l)
}
