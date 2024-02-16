package jwt

import (
	"errors"
	"fmt"
)

// [ Static errors ]

// Secret key errors
var (
	ErrInvalidSecret = errors.New("jwt: the secret key was not the one used to sign")
)

// Token errors
var (
	ErrInvalidToken = errors.New("jwt: token is not valid")
)

// [ Dynamic errors ]

type errShortSecret struct {
	length int // min length of secret
}

func (e *errShortSecret) Error() string {
	return fmt.Sprintf("jwt: the secret key is too short, min length: %d", e.length)
}

func isShortSecretError(err error) bool {
	_, ok := err.(*errShortSecret)
	return ok
}
