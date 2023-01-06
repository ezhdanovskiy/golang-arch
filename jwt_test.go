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

	keys, err := NewKeys()
	require.NoError(t, err)

	token, err := createToken(c, keys.CurrentKid(), keys.CurrentKey())
	require.NoError(t, err)

	claims, err := parseToken(token, keys)
	require.NoError(t, err)

	assert.EqualValues(t, 5, claims.SessionID)
}

func TestJWT_RotatingKeys(t *testing.T) {
	keys, err := NewKeys()
	require.NoError(t, err)
	key1 := keys.CurrentKey()

	c1 := &UserClaims{SessionID: 1}
	c1.ExpiresAt = jwt.NewNumericDate(time.Now().AddDate(0, 0, 1))
	token1, err := createToken(c1, keys.CurrentKid(), keys.CurrentKey())
	require.NoError(t, err)

	// RotatingKeys
	require.NoError(t, keys.GenerateNewKey())
	key2 := keys.CurrentKey()
	require.NotEqual(t, key1, key2)

	c2 := &UserClaims{SessionID: 2}
	c2.ExpiresAt = jwt.NewNumericDate(time.Now().AddDate(0, 0, 1))
	token2, err := createToken(c2, keys.CurrentKid(), keys.CurrentKey())
	require.NoError(t, err)

	claims1, err := parseToken(token1, keys)
	require.NoError(t, err)
	assert.EqualValues(t, 1, claims1.SessionID)

	claims2, err := parseToken(token2, keys)
	require.NoError(t, err)
	assert.EqualValues(t, 2, claims2.SessionID)
}
