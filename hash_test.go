package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	hashedPass1, err := hashPassword("12345")
	require.NoError(t, err)

	err = comparePasswords("12345", hashedPass1)
	require.NoError(t, err)

	err = comparePasswords("123456", hashedPass1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not equal")
}
