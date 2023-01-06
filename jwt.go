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

func createToken(c *UserClaims, currentKid string, currentKey *Key) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	t.Header["kid"] = currentKid
	signedToken, err := t.SignedString(currentKey.key)
	if err != nil {
		return "", fmt.Errorf("failed to sign string: %w", err)
	}

	return signedToken, nil
}

func parseToken(signedToken string, keys *Keys) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("wrong signing algorithm")
		}

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}

		k := keys.Key(kid)
		if k == nil {
			return nil, fmt.Errorf("failed to find key by key ID")
		}

		return k.key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("token is not valid")
	}
	return t.Claims.(*UserClaims), nil
}
