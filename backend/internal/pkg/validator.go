package pkg

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return map[string]string{
			"requrest":"invalid request",
		}
	}

	messages := make(map[string]string)

	for _, fieldErr := range validationErrors {
		field := fieldErr.Field()

		switch field {
			case "Name" : 
				messages["name"] = validationMessage("name", fieldErr.Tag())
			case "Code" : 
				messages["code"] = validationMessage("code", fieldErr.Tag())
			case "Email":
				messages["email"] = validationMessage("email", fieldErr.Tag())
			case "Password" :
				messages["password"] = validationMessage("password", fieldErr.Tag())

			default : 
				messages[field] = "invalid value"
			}
	}

	return messages
}

func validationMessage(field, tag string) string {
    switch tag {
    case "required":
        return field + " is required"

    case "email":
        return field + " must be a valid email"

    default:
        return field + " is invalid"
    }
}