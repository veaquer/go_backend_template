package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func registerCustomRules(v *validator.Validate) {
	v.RegisterValidation("username", validateUsername)
	v.RegisterValidation("password", validatePassword)
}

func validateUsername(fl validator.FieldLevel) bool {
	// Alphanumeric + underscore, 3-30 chars
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`)
	return re.MatchString(fl.Field().String())
}

func validatePassword(fl validator.FieldLevel) bool {
	pwd := fl.Field().String()

	if len(pwd) < 8 {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(pwd)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pwd)
	hasDigit := regexp.MustCompile(`\d`).MatchString(pwd)

	return hasLower && hasUpper && hasDigit
}
