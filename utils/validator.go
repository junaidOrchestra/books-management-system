package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// ✅ Generic Validation Function for Any Struct
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
