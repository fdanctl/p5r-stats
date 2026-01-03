package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func ValidateParam(value any, name string) error {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if v.Len() == 0 {
			return fmt.Errorf("Field '%s' is required", name)
		}
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return fmt.Errorf("Field '%s' is required", name)
		}
	default:
		return errors.New("unsupported type in ValidateParam")
	}

	return nil
}
