package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func DecodeRequestBody[T any](body io.ReadCloser) (T, error) {
	decoder := json.NewDecoder(body)
	var result T
	err := decoder.Decode(&result)
	if err != nil {
		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return result, fmt.Errorf("Malformed JSON at byte %d", syntaxErr.Offset)
		}

		var unmarshalTypeErr *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeErr) {
			return result, fmt.Errorf("Field '%s' has wrong type", unmarshalTypeErr.Field)
		}

		return result, fmt.Errorf("Invalid request result")
	}
	return result, nil
}
