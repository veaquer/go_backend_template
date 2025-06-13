package validator

import (
	"backend_template/pkg/errors/apperror"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	registerCustomRules(validate)
}

func ValidateStruct[T any](s T) error {
	if err := validate.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return apperror.Wrap("Invalid validation", 500, err)
		}

		var sb strings.Builder
		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("Field '%s' failed on the '%s' rule;", err.Field(), err.Tag()))
		}
		return apperror.New(strings.TrimSpace(sb.String()), http.StatusBadRequest)
	}

	return nil
}
