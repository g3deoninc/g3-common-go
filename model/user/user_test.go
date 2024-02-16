package user_test

import (
	"strings"
	"testing"
	"time"

	"github.com/g3deoninc/g3-common-go/model/user"
	"github.com/g3deoninc/g3-common-go/utils/hash"
)

func TestValidate(t *testing.T) {
	hash, err := hash.Hash("password")
	if err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}
	tests := []struct {
		name     string
		u        *user.User
		fieldErr string
	}{
		{
			name: "success",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				AvatarURL:    "https://cdn.g3deon.com/images/avatars/users/5fa3c56c7384a732b018bd46.jpg",
				ProfileColor: "#00FF00",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				Country: "US",
				City:    "San Francisco",

				XP: 999,

				BirthDate:   time.Now(),
				CreatedDate: time.Now(),
				UpdatedDate: time.Now(),

				Password: hash,
			},
		},
		{
			name: "fail_id",
			u: &user.User{
				ID: "not_a_object_id",
			},
			fieldErr: "ID",
		},
		{
			name: "faild_ip",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "not_a_ip",
			},
			fieldErr: "IpAddress",
		},
		{
			name: "fail_username",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "#invalid username",
			},
			fieldErr: "Username",
		},
		{
			name: "fail_email",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "not_an_email",
			},
			fieldErr: "Email",
		},
		{
			name: "fail_display_name",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "validuser",
				Email:    "user@example.com",

				DisplayName: "too_large_invalid_display_name_lol",
			},
			fieldErr: "DisplayName",
		},
		{
			name: "fail_bio",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bhasd46",
				IpAddress: "127.0.0.1",

				Username: "validuser",
				Email:    "user@example.com",

				DisplayName: "too_large_invalid_display_name",
				Bio:         "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Atque laudantium eligendi culpa quod autem beatae ducimus molestiae enim quas, ipsa hic dolore nemo consequuntur delectus quidem, earum cum velit ut obcaecati dolores facilis? Aspernatur quis libero odit, quaerat rem necessitatibus.",
			},
			fieldErr: "Bio",
		},
		{
			name: "fail_avatar_url",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				AvatarURL: "https://invalid.dev/invalid.png",
			},
			fieldErr: "AvatarURL",
		},
		{
			name: "fail_profile_color",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				AvatarURL:    "https://invalid.dev/invalid.png",
				ProfileColor: "not_a_hex",
			},
			fieldErr: "ProfileColor",
		},
		{
			name: "fail_country",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				AvatarURL:    "https://invalid.dev/invalid.png",
				ProfileColor: "#00FF00",

				Country: "LOL",
			},
			fieldErr: "Country",
		},
		{
			name: "fail_city",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				AvatarURL:    "https://invalid.dev/invalid.png",
				ProfileColor: "#00FF00",

				Country: "LOL",
				City:    "",
			},
			fieldErr: "City",
		},
		{
			name: "fail_xp",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				AvatarURL:    "https://invalid.dev/invalid.png",
				ProfileColor: "#00FF00",

				Country: "US",
				City:    "San Francisco",

				XP: -1,
			},
			fieldErr: "XP",
		},
		{
			name: "password_short",
			u: &user.User{
				ID:        "5fa3c56c7384a732b018bd46",
				IpAddress: "127.0.0.1",

				Username: "jhon_doe",
				Email:    "user@example.com",

				DisplayName: "Valid User",
				Bio:         "Short bio",

				AvatarURL:    "https://invalid.dev/invalid.png",
				ProfileColor: "#00FF00",

				Country: "US",
				City:    "San Francisco",

				XP: 999,

				BirthDate:   time.Now(),
				CreatedDate: time.Now(),
				UpdatedDate: time.Now(),

				Password: "this_is_not_a_hash_sha256",
			},
			fieldErr: "Password",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				err     = test.u.Validate()
				wantErr = test.fieldErr != ""
				inErr   = wantErr && strings.Contains(err.Error(), test.fieldErr)
			)

			if !wantErr && err != nil {
				t.Errorf("%s: unexpected error: %v", test.name, err)
			}

			if wantErr && (err == nil || !inErr) {
				t.Errorf("%s: expected error containing '%s', but got: %v", test.name, test.fieldErr, err)
			}
		})
	}
}
