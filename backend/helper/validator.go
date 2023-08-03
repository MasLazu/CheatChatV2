package helper

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Validation []string `json:"validation"`
}

func Validate(writer http.ResponseWriter, data interface{}) error {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		var validationErrorsString []string
		validationErrors := err.(validator.ValidationErrors)

		for _, e := range validationErrors {
			validationErrorsString = append(validationErrorsString, e.Field()+" must be "+e.Tag())
		}

		errorsResponse := ValidationErrorResponse{Validation: validationErrorsString}
		WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", errorsResponse)

		return errors.New("validation errors")
	}
	return nil
}
