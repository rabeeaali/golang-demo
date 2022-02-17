package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

//CheckCodeValidate check code validation
func CheckCodeValidate(code string) error {
	return validation.Errors{
		"code": validation.Validate(code, validation.Required),
	}.Filter()
}
