package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateCustomAlpha(fl validator.FieldLevel) bool {
	var (
		fliedStrng = fl.Field().String()
		regex      = regexp.MustCompile("^[a-zA-Z0-9._]*$")
	)
	return regex.MatchString(fliedStrng)
}
