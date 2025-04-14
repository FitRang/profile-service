package apperror

import (
	"encoding/json"
	"errors"
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

var (
	customErrors = map[string]error{
		"id.required":           errors.New("is required"),
		"id.uuid":               errors.New("has to be uuid"),
		"email.required":        errors.New("is required"),
		"email.email":           errors.New("has to be email"),
		"full_name.required":    errors.New("is required"),
		"phone_number.required": errors.New("is required"),
		"phone_number.e164":     errors.New("has to be in E.164 format"),
	}
)

func CustomValidationError(sourceStruct any, err error) []map[string]string {
	errs := make([]map[string]string, 0)
	switch errTypes := err.(type) {
	case validator.ValidationErrors:
		for _, e := range errTypes {
			errorMap := make(map[string]string)

			key := e.Field() + "." + e.Tag()

			if v, ok := customErrors[key]; ok {
				errorMap[e.Field()] = v.Error()
			} else {
				errorMap[e.Field()] = fmt.Sprintf("custom message is not available: %v", err)
			}
			errs = append(errs, errorMap)
		}
		return errs

	case *json.UnmarshalTypeError:
		errs = append(errs, map[string]string{errTypes.Field: fmt.Sprintf("%v cannot be %v", errTypes.Field, errTypes.Value)})
		return errs
	}
	errs = append(errs, map[string]string{"unknown": fmt.Sprintf("unsupported custom error for: %v", err)})
	return errs
}
