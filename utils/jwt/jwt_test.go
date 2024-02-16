package jwt_test

import (
	"testing"
	"time"

	"github.com/g3deoninc/g3-common-go/utils/jwt"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		data    interface{}
		exp     time.Duration
		secret  string
		wantErr bool
	}{
		{
			name:    "foo/bar",
			data:    map[string]string{"foo": "bar"},
			exp:     10 * time.Hour,
			secret:  "secret_foo",
			wantErr: false,
		},
		{
			name:    "jhonDoe",
			data:    map[string]string{"username": "jhon", "role": "tester"},
			exp:     10 * time.Second,
			secret:  "secret_jhon",
			wantErr: false,
		},
		{
			name:    "short_secret",
			data:    map[string]string{"data": "troll"},
			exp:     1 * time.Microsecond,
			secret:  "short",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token, err := jwt.Tokenize(test.data, test.secret, test.exp)

			if test.wantErr && err == nil {
				t.Errorf("%s: Expected an error, but got none", test.name)
			}

			if !test.wantErr && err != nil {
				t.Errorf("%s: Unexpected error: %v", test.name, err)
			}

			if !test.wantErr {
				if token == "" {
					t.Errorf("%s: Expected a valid token, but got an empty string, %v", test.name, err)
				}
			}
		})
	}
}
