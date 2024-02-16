package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	data                 interface{} // data for the token, Eg. userId: 1
	jwt.RegisteredClaims             // standard claims see in  https://pkg.go.dev/github.com/golang-jwt/jwt/v5@v5.2.0#RegisteredClaims
}

// Tokenize create a new JWT with data and signature
func Tokenize(data interface{}, secret string, exp time.Duration) (string, error) {
	const minLen = 8
	var (
		key    = []byte(secret)
		method = jwt.SigningMethodHS256
		expire = time.Now().Add(exp)
	)

	if len(secret) < minLen {
		return "", &errShortSecret{
			length: minLen,
		}
	}

	claims := tokenClaims{
		data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			Issuer:    "Gedeon Inc.",
		},
	}

	token := jwt.NewWithClaims(method, claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify check signature of a token
func Verify(token, secret string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return ErrInvalidSecret
		}
		return err
	}

	if !parsedToken.Valid {
		return ErrInvalidToken
	}

	return nil
}
