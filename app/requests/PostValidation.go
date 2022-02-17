package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func PostStoreValidate(title string, description string) error {
	return validation.Errors{
		"title":       validation.Validate(title, validation.Required),
		"description": validation.Validate(description, validation.Required, validation.Length(5, 100)),
	}.Filter()
}

func PostUpdateValidate(title string, description string) error {
	return validation.Errors{
		"title":       validation.Validate(title, validation.Required),
		"description": validation.Validate(description, validation.Required, validation.Length(5, 100)),
	}.Filter()
}
