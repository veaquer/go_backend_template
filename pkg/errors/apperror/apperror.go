package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Internal  error  `json:"-"`
	Operation string `json:"operation,omitempty"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("AppError: %s | Internal: %v", e.Message, e.Internal)
	}

	return fmt.Sprintf("AppError: %s", e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Internal
}

// Factory funcs

func New(msg string, code int) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
	}
}

func Wrap(msg string, code int, internal error) *AppError {
	return &AppError{
		Message:  msg,
		Code:     code,
		Internal: internal,
	}
}

// Common helpers

var (
	ErrUnauthorized = New("unauthorized", http.StatusUnauthorized)
	ErrForbidden    = New("forbidden", http.StatusForbidden)
	ErrNotFound     = New("resource not found", http.StatusNotFound)
	ErrInternal     = New("internal server error", http.StatusInternalServerError)
	ErrInvalidInput = New("invalid input", http.StatusBadRequest)
)
