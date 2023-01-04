package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	const msg = "message0"
	signature, err := signMsg([]byte(msg))
	require.NoError(t, err)

	same1, err := checkSign([]byte(msg), signature)
	require.NoError(t, err)
	assert.True(t, same1)

	same2, err := checkSign([]byte(msg+"123"), signature)
	require.NoError(t, err)
	assert.False(t, same2)
}
