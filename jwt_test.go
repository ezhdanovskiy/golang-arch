package main

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {
	c := &UserClaims{SessionID: 5}
	c.ExpiresAt = jwt.NewNumericDate(time.Now().AddDate(0, 0, 1))

	token, err := createToken(c)
	require.NoError(t, err)

	claims, err := parseToken(token)
	require.NoError(t, err)

	assert.EqualValues(t, 5, claims.SessionID)
}
