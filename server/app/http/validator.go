package http

import validator "gopkg.in/go-playground/validator.v9"

// CustomValidator validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
