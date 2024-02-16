package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
# Create a hash string with sha256 algorithm

Configuration:

  - Min Length: 8
  - Max Length: 64
*/
func Hash(password string) (string, error) {
	const (
		minLength = 8
		maxLength = 64
	)

	if len(password) < minLength || len(password) > maxLength {
		return "", &errLengthPassword{
			minLength: minLength,
			maxLength: maxLength,
		}
	}

	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	return hashedPassword, nil
}

// Compare 2 hashed strings
func Compare(password, hashed string) error {
	if hashed == password {
		return ErrSameHashAndPassword
	}

	if hashed == "" {
		return ErrEmptyPassword
	}

	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	if hashed != hashedPassword {
		return ErrHashMismatch
	}

	return nil
}
