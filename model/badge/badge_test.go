package badge_test

import (
	"strings"
	"testing"
	"time"

	"github.com/g3deoninc/g3-common-go/model/badge"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		b        *badge.Badge
		fieldErr string
	}{
		{
			name: "success",
			b: &badge.Badge{
				ID:          "60f0cde0b33aeb001f5991a7",
				IconUrl:     "https://cdn.g3deon.com/badge",
				IconCode:    "G3_TEST_GO",
				UserCount:   100,
				CreatedDate: time.Now(),
			},
		},
		{
			name: "faild_id",
			b: &badge.Badge{
				ID: "",
			},
			fieldErr: "ID",
		},
		{
			name: "faild_icon_code",
			b: &badge.Badge{
				ID:       "60f0cde0b33aeb001f5991a7",
				IconCode: "invalid_code",
			},
			fieldErr: "IconCode",
		},
		{
			name: "faild_icon_url",
			b: &badge.Badge{
				ID:       "609c8ac6a49d287bb80d6f3d",
				IconCode: "G3_TEST_CODE",
				IconUrl:  "invalid_url",
			},
			fieldErr: "IconUrl",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				err     = test.b.Validate()
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
