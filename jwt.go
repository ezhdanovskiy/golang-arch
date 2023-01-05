package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.RegisteredClaims
	SessionID int64
}

func (c *UserClaims) Valid() error {
	if !c.VerifyExpiresAt(time.Now(), true) {
		return fmt.Errorf("token has expired")
	}
	if c.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}

	return nil
}

const jwtSignedKey = "key"

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString([]byte(jwtSignedKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign string: %w", err)
	}

	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("wrong signing algorithm")
		}
		return []byte(jwtSignedKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("token is not valid")
	}
	return t.Claims.(*UserClaims), nil
}
