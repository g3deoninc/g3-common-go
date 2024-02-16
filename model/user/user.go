package user

import (
	"errors"
	"time"

	"github.com/g3deoninc/g3-common-go/utils/validation"
	"github.com/go-playground/validator/v10"
)

var (
	ErrPasswordToShort = errors.New("password must be at least 8 characters long")
	ErrEmptyPassword   = errors.New("password must not be empty")
)

// [ Structs ]
type User struct {
	ID        string `json:"id" bson:"_id" validate:"required,mongodb"`
	IpAddress string `json:"ip_address" bson:"_ip_address" validate:"required,ip"`

	Username string `json:"username" bson:"_id" validate:"required,min=2,max=32,g3_alpha,lowercase"`
	Email    string `json:"email" bson:"email" validate:"required,email"`

	DisplayName string `json:"display_name" bson:"display_name" validate:"required,min=1,max=32"`
	Bio         string `json:"bio" bson:"bio" validate:"required,max=264"`

	AvatarURL    string `json:"avatar_url" bson:"avatar_url" validate:"required,url,g3_cdn"`
	ProfileColor string `json:"banner_url" bson:"banner_url" validate:"required,hexcolor"`

	Country string `json:"country" bson:"country" validate:"required,iso3166_1_alpha2"`
	City    string `json:"city" bson:"city" validate:"required,min=1,max=264"`

	XP int `json:"level" bson:"level" validate:"required,min=0"`
	// Badges ObjectId's strings
	Badges []string `json:"badges" bson:"badges"`

	BirthDate   time.Time `json:"birthday_date" bson:"required, birthday_date" validate:"required"`
	CreatedDate time.Time `json:"created_at" bson:"required,created_at" validate:"required"`
	UpdatedDate time.Time `json:"updated_at" bson:"required,updated_at" validate:"required"`

	// SHA256 Hash string
	Password string `json:"password" bson:"password" validate:"required,sha256"`
}

// [ Methods ]
func (u *User) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validation.RegisterCustomValidators(*validate); err != nil {
		return err
	}

	return validate.Struct(u)
}
