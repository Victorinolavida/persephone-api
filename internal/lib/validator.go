package lib

import (
	"errors"
	"github.com/Victorinolavida/persephone-api/internal/adapters/dto"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validator struct {
	*validator.Validate
}

var (
	ErrPasswordMismatch = errors.New("passwords do not match")
)

func NewValidator() *Validator {
	v := &Validator{
		validator.New(validator.WithRequiredStructEnabled()),
	}
	return v
}

func (v *Validator) ValidateStruct(data interface{}) error {
	err := v.Validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func (v *Validator) UserPassword(user dto.User) error {
	if strings.Compare(user.PasswordConfirmation, user.Password) != 0 {
		return ErrPasswordMismatch
	}
	return nil
}
