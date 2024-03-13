package helper

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) []Error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]Error, len(ve))
		for i, fe := range ve {
			out[i] = Error{fe.Field(), getErrorMsg(fe)}
		}
		return out
	}
	return []Error{{Msg: "unknown error"}}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "alpha":
		return "Should be alphabetical"
	case "startswith":
		return "Should be start with " + fe.Param()
	case "email":
		return "Should be email account"
	}
	return "Unknown error"
}
