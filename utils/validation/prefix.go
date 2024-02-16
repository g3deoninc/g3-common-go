package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateBadgePrefix(fl validator.FieldLevel) bool {
	const badgePrefix string = "G3_"
	var field string = fl.Field().String()
	return strings.HasPrefix(field, badgePrefix)
}

func ValidateCDNPrefix(fl validator.FieldLevel) bool {
	var fliedStrng = fl.Field().String()
	const cdnPrefix = "https://cdn.g3deon.com/"
	return strings.HasPrefix(fliedStrng, cdnPrefix)
}
