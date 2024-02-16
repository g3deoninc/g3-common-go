package hash_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/g3deoninc/g3-common-go/utils/hash"
)

const (
	testIterations = 10
)

func TestHash(t *testing.T) {
	for i := 0; i < testIterations; i++ {
		t.Run("password", func(t *testing.T) {
			randomPassword := generateRandomPassword()
			hashedPassword, err := hash.Hash(randomPassword)
			if err != nil {
				t.Errorf("Error hashing password: %v", err)
			}

			if len(hashedPassword) == 0 {
				t.Error("Hashed password should not be empty")
			}
		})
	}
}

func TestCompare(t *testing.T) {
	for i := 0; i < testIterations; i++ {
		t.Run("compare", func(t *testing.T) {
			randomPassword := generateRandomPassword()
			hashedPassword, err := hash.Hash(randomPassword)
			if err != nil {
				t.Errorf("Error hashing password: %v", err)
			}

			err = hash.Compare(randomPassword, hashedPassword)
			if err != nil {
				t.Errorf("Error comparing password with itself: %v", err)
			}

			err = hash.Compare("otracontraseña", hashedPassword)
			if err == nil {
				t.Error("Comparison with different password should result in an error")
			}
		})
	}
}

func generateRandomPassword() string {
	const (
		minLength = 8
		maxLength = 20
	)

	ramdon := rand.New(rand.NewSource(time.Now().UnixNano()))
	passwordLength := ramdon.Intn(maxLength-minLength+1) + minLength
	passwordChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	password := make([]byte, passwordLength)

	for i := range password {
		password[i] = passwordChars[ramdon.Intn(len(passwordChars))]
	}

	return string(password)
}
