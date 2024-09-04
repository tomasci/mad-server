package app_middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"mad_backend_v1/utils/response"
	"mad_backend_v1/utils/validation"
	"net/http"
)

type requestBodyKey int

const bodyKey requestBodyKey = iota

func validateStruct[T any](data T) validation.ResultValidationErrors {
	result := make(validation.ResultValidationErrors)
	validate := validation.Validate()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldKey := validation.FindJsonTagName(data, err.Field())
			result[fieldKey] = append(result[fieldKey], err.Tag())
		}
	}

	return result
}

func RequestBodyMiddleware[T any]() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			var body T

			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				response.Error[any](writer, 400, nil, errors.New("invalid_request_body"))
				return
			}

			// default validation
			validationErrors := validateStruct(body)
			if len(validationErrors) > 0 {
				response.Error(writer, 400, validationErrors, errors.New("validation_failed"))
				return
			}

			// additional validation, just in case...
			if additionalValidator, ok := any(body).(validation.Validator); ok {
				validationErrors := additionalValidator.Validate()
				if len(validationErrors) > 0 {
					response.Error(writer, 400, validationErrors, errors.New("validation_failed"))
					return
				}
			}

			ctx := context.WithValue(request.Context(), bodyKey, body)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}

func GetRequestBody[T any](ctx context.Context) T {
	body, ok := ctx.Value(bodyKey).(T)

	if !ok {
		var zero T
		return zero
	}

	return body
}
