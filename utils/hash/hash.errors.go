package hash

import (
	"errors"
	"fmt"
)

var (
	ErrSameHashAndPassword = errors.New("password and the hashed password cannot be the same")
	ErrEmptyPassword       = errors.New("hashed password or password must be not empty")
	ErrHashMismatch        = errors.New("hashed password and password must match")
)

type errLengthPassword struct {
	minLength int // min length of password
	maxLength int // max length of password
}

func (e *errLengthPassword) Error() string {
	return fmt.Sprintf(
		"the password length is invalid, the minimum length is (%d) and the maximum length is (%d)",
		e.minLength,
		e.maxLength,
	)
}

func IsLengthPasswordError(err error) bool {
	_, ok := err.(*errLengthPassword)
	return ok
}
