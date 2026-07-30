package apperr

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Code string

const (
	CodeNotFound   Code = "not_found"
	CodeValidation Code = "validation_error"
	CodeForbidden  Code = "forbidden"
	CodeConflict   Code = "conflict"
	CodeInternal   Code = "internal_error"
)

type AppError struct {
	Code     Code
	Message  string // safe to show to the client
	Status   int
	Internal bool  // true = don't leak details, log server-side only
	Err      error // wrapped original error, for logging
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Constructors for common cases

func NotFound(resource string) *AppError {
	return &AppError{
		Code:    CodeNotFound,
		Message: resource + " not found",
		Status:  fiber.StatusNotFound,
	}
}

func Validation(code Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  fiber.StatusBadRequest,
	}
}

func Forbidden(message string) *AppError {
	return &AppError{
		Code:    CodeForbidden,
		Message: message,
		Status:  fiber.StatusForbidden,
	}
}
func Conflict(message string) *AppError {
	return &AppError{
		Code:    CodeConflict,
		Message: message,
		Status:  fiber.StatusConflict,
	}
}

func Internal(err error) *AppError {
	return &AppError{
		Code:     CodeInternal,
		Message:  "internal server error",
		Status:   http.StatusInternalServerError,
		Internal: true,
		Err:      err,
	}
}
