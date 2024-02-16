package validation

import "github.com/go-playground/validator/v10"

/*
# Register all custom validation functions.

Available tags:

  - `g3_badge`: Validate badge prefix, e.g., "G3_BADGE".
  - `g3_cdn`: Validate a g3doen CDN prefix, e.g., "cdn.g3deon.com/...".
  - `g3_alpha`: Validate alphanumeric characters with characters "_, .".
*/
func RegisterCustomValidators(validate validator.Validate) error {
	if err := validate.RegisterValidation("g3_badge", ValidateBadgePrefix); err != nil {
		return err
	}

	if err := validate.RegisterValidation("g3_cdn", ValidateCDNPrefix); err != nil {
		return err
	}

	if err := validate.RegisterValidation("g3_alpha", ValidateCustomAlpha); err != nil {
		return err
	}
	return nil
}
