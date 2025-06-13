package apperror

import (
	"fmt"
	"net/http"
)

type ConflictError struct {
	Field string
	Msg   string
}

func (c *ConflictError) Error() string {
	return fmt.Sprintf("conflict on %s : %s", c.Field, c.Msg)
}

func (c *ConflictError) GetCode() int {
	return http.StatusConflict
}

func (c *ConflictError) GetMessage() string {
	return c.Msg
}

func NewConflict(Field, Msg string) error {
	return &ConflictError{Msg: Msg, Field: Field}
}
