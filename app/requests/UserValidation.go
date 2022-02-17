package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pkg/errors"
)

func RegisterValidate(name string, email string, password string) error {
	return validation.Errors{
		"name":     validation.Validate(name, validation.Required, validation.Length(5, 190)),
		"email":    validation.Validate(email, validation.Required, is.Email),
		"password": validation.Validate(password, validation.Required, validation.Length(6, 100)),
	}.Filter()
}

func LoginValidate(email string, password string) error {
	return validation.Errors{
		"email":    validation.Validate(email, validation.Required, is.Email),
		"password": validation.Validate(password, validation.Required, validation.Length(5, 100)),
	}.Filter()
}

func ForgetPasswordValidate(email string) error {
	return validation.Errors{
		"email": validation.Validate(email, validation.Required, is.Email),
	}.Filter()
}

func ResetPasswordValidate(email string, password string, passwordConfirmation string) error {
	return validation.Errors{
		"email": validation.Validate(email, validation.Required, is.Email),
		"password": validation.Validate(
			password,
			validation.Required,
			validation.Length(5, 100),
			validation.By(passwordEquals(passwordConfirmation)),
		),
	}.Filter()
}

func passwordEquals(str string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != str {
			return errors.New("passwords do not match")
		}
		return nil
	}
}
