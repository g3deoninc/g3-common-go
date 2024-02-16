package badge

import (
	"time"

	"github.com/g3deoninc/g3-common-go/utils/validation"
	"github.com/go-playground/validator/v10"
)

// [ Structs ]
type Badge struct {
	ID string `json:"id" bson:"_id" validate:"required,mongodb"`

	IconCode string `json:"icon_code" bson:"icon_code" validate:"required,min=4,max=32,g3_badge"`
	IconUrl  string `json:"icon_url" bson:"icon_url" validate:"required,url,g3_cdn"`

	UserCount int `json:"user_count" bson:"user_count" validate:"required,min=0"`

	CreatedDate time.Time `json:"created_at" bson:"created_at"`
}

// [ Methods ]
func (b *Badge) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validation.RegisterCustomValidators(*validate); err != nil {
		return err
	}

	return validate.Struct(b)
}
