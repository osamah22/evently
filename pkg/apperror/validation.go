package apperr

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FromBindErr converts an error returned by Fiber's request binding
// (c.Bind().Body(...)) into a client-safe validation AppError. It covers both
// struct validation failures (validator.ValidationErrors) and body parse
// failures (e.g. malformed JSON).
func FromBindErr(err error) *AppError {
	if verrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
		msgs := make([]string, 0, len(verrs))
		for _, fe := range verrs {
			msgs = append(msgs, fmt.Sprintf("%s failed validation on %q", fe.Field(), fe.Tag()))
		}
		return Validation(CodeValidation, strings.Join(msgs, "; "))
	}

	return Validation(CodeValidation, "invalid request body")
}
