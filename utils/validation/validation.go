package validation

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Validator interface {
	Validate() ResultValidationErrors
}

type ResultValidationErrors map[string][]string

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

// Validate is used to return one and only instance of validator
func Validate() *validator.Validate {
	return validate
}

// FindJsonTagName taken from https://github.com/go-playground/validator/issues/877
func FindJsonTagName(i interface{}, original string) string {
	reflected := reflect.ValueOf(i)

	switch reflected.Kind() {
	case reflect.Ptr:
		reflected = reflected.Elem()
	case reflect.Struct:
		break
	default:
		// Fields that are not structs will be ignored
		return original
	}

	// Find the field from the struct and extract tag
	for i := 0; i < reflected.Type().NumField(); i++ {
		f := reflected.Type().Field(i)
		if f.Name == original {
			tag := f.Tag.Get("json")
			if tag != "" {
				return tag
			} else {
				return original
			}
		}
	}

	// Always return the original field name
	// We should never get here. If we do, it means the struct being
	// checked is the wrong one. The field should always be found.
	return original
}
