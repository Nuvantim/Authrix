package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// A pre-configured validator instance.
var validate = func() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		if name := f.Tag.Get("json"); name != "" {
			return name
		}
		return f.Name
	})
	return v
}()

// msgForTag returns a human-readable error message for a given validation tag.
var staticErrorMessages = map[string]string{
	"required": "this field is required",
	"email":    "invalid email format",
	"default":  "validation failed",
}

func msgForTag(tag, param string) string {
	if msg, ok := staticErrorMessages[tag]; ok {
		return msg
	}

	switch tag {
	case "gte":
		return fmt.Sprintf("value must be ≥ %s", param)
	case "lte":
		return fmt.Sprintf("value must be ≤ %s", param)
	default:
		return staticErrorMessages["default"]
	}
}

// validate performs validation on a struct.
func Validates[T any](data T) error {
	// Validate using the go-playground validator.
	if err := validate.Struct(data); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) && len(ve) > 0 {
			fe := ve[0]
			return fmt.Errorf("param(%s): %s", fe.Field(), msgForTag(fe.Tag(), fe.Param()))
		}
	}
	return nil
}
